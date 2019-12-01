# RTSPtoWSMP4f

RTSP Stream to WebBrowser MSE over WebSocket based MP4f segments

![RTSPtoWSMP4f image](doc/demo4.png)

## Team

Deepch - https://github.com/deepch streaming developer

Dmitry - https://github.com/vdalex25 web developer

## Installation
1.
```bash
go get github.com/deepch/RTSPtoWSMP4f
```
2.
```bash
cd src/github.com/deepch/RTSPtoWSMP4f
```
3.
```bash
go run *.go
```
4.
```bash
open web browser http://127.0.0.1:8083
```

## Configuration

### Edit file config.json

format:

```bash
{
  "server": {
    "http_port": ":8083"
  },
  "streams": {
    "demo1": {
      "url": "rtsp://170.93.143.139/rtplive/470011e600ef003a004ee33696235daa"
    },
    "demo2": {
      "url": "rtsp://170.93.143.139/rtplive/470011e600ef003a004ee33696235daa"
    },
    "demo3": {
      "url": "rtsp://170.93.143.139/rtplive/470011e600ef003a004ee33696235daa"
    }
  }
}
```

## Limitations

Video Codecs Supported: H264 all profiles

Audio Codecs Supported: AAC mono

## Test

CPU usage 0.2% one core cpu intel core i7 / stream
