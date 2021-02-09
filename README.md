# RTSPtoWSMP4f

RTSP Stream to WebBrowser MSE over WebSocket based MP4f segments

full native! not use ffmpeg or gstreamer

if you need RTSPtoWebRTC use https://github.com/deepch/RTSPtoWebRTC

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
   "H264_AAC": {
      "url": "rtsp://wowzaec2demo.streamlock.net/vod/mp4:BigBuckBunny_115k.mov"
    }
  }
}
```

## Limitations

Video Codecs Supported: H264 all profiles, H265 work only safari and (IE hw video card)

Audio Codecs Supported: AAC

## Test

CPU usage 0.2% one core cpu intel core i7 / stream



[![paypal.me/AndreySemochkin](https://ionicabizau.github.io/badges/paypal.svg)](https://www.paypal.me/AndreySemochkin) - You can make one-time donations via PayPal. I'll probably buy a ~~coffee~~ tea. :tea:


