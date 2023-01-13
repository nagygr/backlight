# Backlight

Application that sets the backlight brightness on Linux machines.

## Usage

The application uses the commands (`brightness`, `actual_brightness`,
`max_brightness`) under `/sys/class/backlight/<device>`. The `device` part is
specific to the video card in the host machine. The app tries to determine the
name automatically but it can also be expicitly given using the `-cmd` flag.

The application expects a percentage with which it should increase or decrease
the brightness ([-100 : 100]). It also prints the current and maximum values.
If no percentage is given it only prints the values.

```
backlight [-cmd=<brightness commands directory>] [-p=<percentage>]
```

## Installation

The application can be installed with the following command:

```
go install github.com/nagygr/backlight/cmd/backlight@latest
```

Please note, that this requires Go to be available on the system. It can be
installed from the package manager, e.g.:

```bash
pacman -S go            # Arch Linux, Manjaro, etc.
dnf install golang      # Fedora, Red Hat, etc.
apt-get install golang  # Debian, Ubuntu, etc.
```
