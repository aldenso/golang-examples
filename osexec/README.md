osexec
=====================

Just to remember how to work with os/exec package.

```
[aldenso@golang go]$ go run osexec.go
Running Command: ls, with args: [-la /home/aldenso]
total 24
drwx------. 3 aldenso aldenso 4096 Sep 24 01:10 .
drwxr-xr-x. 3 root root   17 Aug 12  2015 ..
-rw-------. 1 aldenso aldenso   81 Sep 24 01:10 .bash_history
-rw-r--r--. 1 aldenso aldenso   18 Jun 10  2014 .bash_logout
-rw-r--r--. 1 aldenso aldenso  193 Jun 10  2014 .bash_profile
-rw-r--r--. 1 aldenso aldenso  275 Sep 24 01:09 .bashrc
drwxrwxr-x. 2 aldenso aldenso   22 Sep 24 01:08 go
-rw-------. 1 aldenso aldenso  624 Sep 24 01:09 .viminfo

Running Command: date, with args: []
Sat Sep 24 01:10:53 VET 2016

Running Command: df, with args: [-h]
Filesystem               Size  Used Avail Use% Mounted on
/dev/mapper/centos-root  6.7G  1.6G  5.2G  23% /
devtmpfs                 487M     0  487M   0% /dev
tmpfs                    497M     0  497M   0% /dev/shm
tmpfs                    497M  6.6M  490M   2% /run
tmpfs                    497M     0  497M   0% /sys/fs/cgroup
/dev/sda1                497M  189M  309M  38% /boot
tmpfs                    100M     0  100M   0% /run/user/0
tmpfs                    100M     0  100M   0% /run/user/1000
```
