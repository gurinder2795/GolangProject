# Audio Stream Mixer

This Go application uses GStreamer to mix two audio streams (a test tone and the user's microphone audio). The test tone is output only if the user is not speaking.

## Prerequisites

- Go (Golang) 1.13+
- GStreamer 1.16+

## Installation

1. Install Go (Golang) from [https://golang.org/doc/install](https://golang.org/doc/install)
2. Install GStreamer from [https://gstreamer.freedesktop.org/download/](https://gstreamer.freedesktop.org/download/)

## Setup

1. Clone the repository:

```bash
git clone https://github.com/yourusername/audiostream-mixer.git
cd audiostream-mixer



Install the required Go packages:

go get github.com/notedit/gstreamer-go
go get github.com/notedit/gstreamer-go/gst



Run the application:
go run main.go


Usage
The application generates a 440 Hz test tone and captures audio from the user's microphone.

The test tone is output only if the user is not speaking.

If the user speaks, the test tone is muted.