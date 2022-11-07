wget https://busybox.net/downloads/binaries/1.35.0-x86_64-linux-musl/busybox
chmod a+x busybox
mkdir -p  /tmp/rootfs/proc /tmp/rootfs/dev rootfs/bin
./busybox --install rootfs/bin
mv rootfs/bin /tmp/rootfs
rm -r busybox rootfs
