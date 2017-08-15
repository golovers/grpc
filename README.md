# Getting started with gRPC

This project contains some examples to help you to easier to get started with gRPC. I will try to cover all the cases but if you think I'm missing something, please let me know by dropping me an [email](phamthethanh1088@gmail.com )

## Notice

Before running into the details of each version of the examples, make sure you installed Protocol Buffer by using following commands:

```shell
wget https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip
unzip protoc-3.3.0-linux-x86_64.zip -d protoc
sudo cp protoc/bin/protoc /usr/local/bin
sudo mkdir -p /usr/local/include
sudo cp -R protoc/include/* /usr/local/include/
rm -rf protoc
rm -rf protoc-3.3.0-linux-x86_64.zip
```

**Note**: *In the above commands, we moved the include folder of Protobuf to ```/usr/local/include``` for convenient for later usages, hence if you already installed Protobuf, make sure you do the same to ensure running the examples without any error.*

