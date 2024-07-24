### Custom implementation of QRcode
But with a(several) caveats
- Custom Encoder/Decoder (no one can/should use it lol)
- Uses binary data from the get-go
- No error correction
- Dynamic matrix size


## Usage
```go
$ go run .
```

## TODO
- [ ] Implement error correction
- [ ] Create the decoder
- [ ] Find a way to resize the matrix dynamically
- [ ] A way to store encoding/size of content in the QR code
