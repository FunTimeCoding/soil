package constant

const WirelessSample = `
————————————————————————————————————
NETWORK
————————————————————————————————————
    Primary IPv4             : en0 (Wi-Fi)
————————————————————————————————————
WIFI
————————————————————————————————————
    MAC Address              : aa:bb:cc:dd:ee:ff
    RSSI                     : -45 dBm
    Channel                  : 5g44/80
————————————————————————————————————
AWDL
————————————————————————————————————
    AWDL Enabled             : Yes
    Channel Sequence         : {(44, 80MHz), (6, 20MHz)}
`

const BluetoothSample = `{
  "SPBluetoothDataType": [
    {
      "device_connected": [
        {"Example Mouse": {"device_minorType": "mouse"}}
      ],
      "device_not_connected": [
        {"Example Keyboard": {}}
      ]
    }
  ]
}`
