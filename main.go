package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"

    "github.com/notedit/gstreamer-go/gst"
)

func main() {
    // Initialize GStreamer
    gst.Init(nil)

    // Create the pipeline
    pipeline, err := gst.ParseLaunch(`
        audiotestsrc is-live=true wave=sine freq=440 volume=0.5 ! audioconvert ! volume name=tonemixer ! audiomixer name=mix 
        autoaudiosrc ! audioconvert ! level interval=100000000 ! mix.
        mix. ! autoaudiosink
    `)
    if err != nil {
        fmt.Println("Failed to create pipeline:", err)
        return
    }

    // Get the volume and level elements
    toneMixer, err := pipeline.GetByName("tonemixer")
    if err != nil {
        fmt.Println("Failed to get volume element:", err)
        return
    }
    level, err := pipeline.GetByName("level")
    if err != nil {
        fmt.Println("Failed to get level element:", err)
        return
    }

    // Connect the level element's message handler
    bus := pipeline.GetBus()
    bus.AddSignalWatch()
    bus.Connect("message::element", func(msg *gst.Message) {
        s := msg.GetStructure()
        if s != nil && s.HasName("level") {
            peak := s.GetValueArray("peak")[1].(float64) // Get the peak value of the second channel (user microphone)

            if peak > 0.01 {
                toneMixer.SetProperty("volume", 0.0) // Mute the test tone
            } else {
                toneMixer.SetProperty("volume", 0.5) // Unmute the test tone
            }
        }
    })

    // Start the pipeline
    pipeline.SetState(gst.StatePlaying)

    // Handle Ctrl+C to gracefully exit
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    <-c

    // Stop the pipeline
    pipeline.SetState(gst.StateNull)
}
