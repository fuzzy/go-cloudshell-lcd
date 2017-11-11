# go-cloudshell-lcd

### Installation

* First install the binary:

```
$ go get -v github.com/fuzzy/go-cloudshell-lcd
$ sudo cp ${GOPATH}/bin/go-cloudshell-lcd /usr/bin
```

* Next copy the config file to /etc and edit to taste:

```
$ sudo cp ${GOPATH}/src/github.com/fuzzy/go-cloudshell-lcd/go-cloudshell.yml /etc
```

* Boot time configuration for distros with systemd

Use a unit file such as this one (adapted from https://github.com/mdrjr/cloudshell_lcd/blob/master/cloudshell):

```
[Unit]
Description="ODROID Cloudshell LCD Info"
DefaultDependencies=no
Requires=sysinit.target
After=sysinit.target

[Service]
Type=simple
ExecStart=/bin/cloudshell-lcd

[Install]
WantedBy=basic.target
WantedBy=sysinit.target
```

or an entry in /etc/rc.local or whatever facility your distribution provides executing the wrapper script /bin/cloudshell-lcd. You can use something like the following (adapted from https://github.com/mdrjr/cloudshell_lcd/blob/master/cloudshell)

```
#!/bin/bash
# hardkernel CloudShell Screen update

# Disable LCD Slepp mode
echo -e '\033[9;0]' > /dev/tty1
# console font
# More fonts on: /usr/share/consolefonts
export TERM="linux"
export CONSOLE_FONT="Lat7-Fixed18"
#export CONSOLE_FONT="Lat15-TerminusBold20x10"
# Output Console (ttyX)
export OUTPUT_CONSOLE="1"
oc="/dev/tty$OUTPUT_CONSOLE"

# font setup
setfont $CONSOLE_FONT > $oc

while true; do
  # Ensure that we are in the right TTY
  chvt $OUTPUT_CONSOLE
	/usr/local/golang/path/bin/go-cloudshell-lcd >$oc
done
			
```