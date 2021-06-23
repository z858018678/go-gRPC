###证书生成
####私钥  
    openssl ecparam -genkey -name secp384r1 -out server.key
#####自签公钥
    openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650