# RTSPtoWSMP4f

RTSP Stream to WebBrowser MSE over WebSocket based MP4f segments

full native! not use ffmpeg or gstreamer

if you need RTSPtoWebRTC use https://github.com/deepch/RTSPtoWebRTC

![RTSPtoWSMP4f image](doc/demo4.png)


### Download Source

1. Download source
   ```bash 
   $ git clone https://github.com/deepch/RTSPtoWSMP4f  
   ```
3. CD to Directory
   ```bash
    $ cd RTSPtoWSMP4f/
   ```
4. Test Run
   ```bash
    $ GO111MODULE=on go run *.go
   ```
5. Open Browser
    ```bash
    open web browser http://127.0.0.1:8083 work chrome, safari, firefox
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

## Team

Deepch - https://github.com/deepch streaming developer

Dmitry - https://github.com/vdalex25 web developer

## Other Example

Examples of working with video on golang

- [RTSPtoWeb](https://github.com/deepch/RTSPtoWeb)
- [RTSPtoWebRTC](https://github.com/deepch/RTSPtoWebRTC)
- [RTSPtoWSMP4f](https://github.com/deepch/RTSPtoWSMP4f)
- [RTSPtoImage](https://github.com/deepch/RTSPtoImage)
- [RTSPtoHLS](https://github.com/deepch/RTSPtoHLS)
- [RTSPtoHLSLL](https://github.com/deepch/RTSPtoHLSLL)

[![paypal.me/AndreySemochkin](https://ionicabizau.github.io/badges/paypal.svg)](https://www.paypal.me/AndreySemochkin) - You can make one-time donations via PayPal. I'll probably buy a ~~coffee~~ tea. :tea: