# fccb

A command line tool for building Fedora CoreOS Configs (FCC): <https://docs.fedoraproject.org/en-US/fedora-coreos/producing-ign/>.

### Examples

```shell
$ fccb add file test.fcc mymotd.txt -p /etc/motd # Add motd 
$ fccb add unit test.fcc -f test.service -d test.service.d/ # Add Systemd service unit and some drop-ins
$ fccb add user test.fcc -n bostrt -G floppy -s /bin/ksh # Add user in floppy and uses KSH 
```

### TODO:

- Update the input FCC file instead of just stdout.
- Implement rest of FCC fields in <https://docs.fedoraproject.org/en-US/fedora-coreos/fcct-config/>.
- Clean up code if this ends up be really useful.
