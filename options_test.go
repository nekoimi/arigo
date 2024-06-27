package arigo

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionFormat(t *testing.T) {
	data := []byte(`{
		"allow-overwrite": "false",
		"allow-piece-length-change": "false",
		"always-resume": "true",
		"async-dns": "true"
	}`)

	var options Options
	assert.NoError(t, json.Unmarshal(data, &options), "Couldn't unmarshal JSON")

	assert.Equal(t, Options{
		AllowOverwrite:         false,
		AllowPieceLengthChange: false,
		AlwaysResume:           true,
		AsyncDNS:               true,
	}, options)
}

func TestOptionsAll(t *testing.T) {
	data := []byte(`
		{
        "allow-overwrite": "false",
        "allow-piece-length-change": "true",
        "always-resume": "false",
        "async-dns": "true",
        "auto-file-renaming": "true",
        "auto-save-interval": "20",
        "bt-detach-seed-only": "true",
        "bt-enable-hook-after-hash-check": "true",
        "bt-enable-lpd": "true",
        "bt-force-encryption": "true",
        "bt-hash-check-seed": "true",
        "bt-load-saved-metadata": "true",
        "bt-max-open-files": "100",
        "bt-max-peers": "128",
        "bt-metadata-only": "false",
        "bt-min-crypto-level": "plain",
        "bt-prioritize-piece": "head=32M,tail=32M",
        "bt-remove-unselected-file": "true",
        "bt-request-peer-speed-limit": "10485760",
        "bt-require-crypto": "false",
        "bt-save-metadata": "true",
        "bt-seed-unverified": "false",
        "bt-stop-timeout": "0",
        "bt-tracker": "",
        "bt-tracker-connect-timeout": "10",
        "bt-tracker-interval": "0",
        "bt-tracker-timeout": "10",
        "ca-certificate": "/etc/ssl/certs/ca-certificates.crt",
        "check-certificate": "true",
        "check-integrity": "false",
        "conditional-get": "false",
        "conf-path": "/etc/aria2/aria2.conf",
        "connect-timeout": "10",
        "console-log-level": "notice",
        "content-disposition-default-utf8": "true",
        "continue": "true",
        "daemon": "true",
        "deferred-input": "false",
        "dht-entry-point": "dht.transmissionbt.com:6881",
        "dht-entry-point6": "dht.transmissionbt.com:6881",
        "dht-file-path": "/home/nas-data/.aria2/dht.dat",
        "dht-file-path6": "/home/nas-data/.aria2/dht6.dat",
        "dht-listen-port": "36810",
        "dht-message-timeout": "10",
        "dir": "/mnt/datahub/javtub/0BitDownload",
        "disable-ipv6": "true",
        "disk-cache": "67108864",
        "download-result": "default",
        "dry-run": "false",
        "dscp": "0",
        "enable-color": "true",
        "enable-dht": "true",
        "enable-dht6": "false",
        "enable-http-keep-alive": "true",
        "enable-http-pipelining": "false",
        "enable-mmap": "false",
        "enable-peer-exchange": "true",
        "enable-rpc": "true",
        "event-poll": "epoll",
        "file-allocation": "none",
        "follow-metalink": "true",
        "follow-torrent": "true",
        "force-save": "false",
        "ftp-pasv": "true",
        "ftp-reuse-connection": "true",
        "ftp-type": "binary",
        "hash-check-only": "false",
        "help": "#basic",
        "http-accept-gzip": "true",
        "http-auth-challenge": "false",
        "http-no-cache": "false",
        "human-readable": "true",
        "keep-unfinished-download-result": "true",
        "listen-port": "36810",
        "log": "/var/log/aria2/aria2.log",
        "log-level": "debug",
        "lowest-speed-limit": "0",
        "max-concurrent-downloads": "10",
        "max-connection-per-server": "16",
        "max-download-limit": "0",
        "max-download-result": "1000",
        "max-file-not-found": "10",
        "max-mmap-limit": "9223372036854775807",
        "max-overall-download-limit": "0",
        "max-overall-upload-limit": "2097152",
        "max-resume-failure-tries": "0",
        "max-tries": "0",
        "max-upload-limit": "0",
        "metalink-enable-unique-protocol": "true",
        "metalink-preferred-protocol": "none",
        "min-split-size": "4194304",
        "min-tls-version": "TLSv1.2",
        "netrc-path": "/home/nas-data/.netrc",
        "no-conf": "false",
        "no-file-allocation-limit": "67108864",
        "no-netrc": "true",
        "optimize-concurrent-downloads": "false",
        "parameterized-uri": "false",
        "pause-metadata": "false",
        "peer-agent": "Deluge 1.3.15",
        "peer-id-prefix": "-DE13F0-",
        "piece-length": "1048576",
        "proxy-method": "get",
        "quiet": "true",
        "realtime-chunk-checksum": "true",
        "remote-time": "true",
        "remove-control-file": "false",
        "retry-wait": "10",
        "reuse-uri": "false",
        "rlimit-nofile": "1024",
        "rpc-allow-origin-all": "true",
        "rpc-listen-all": "true",
        "rpc-listen-port": "6800",
        "rpc-max-request-size": "10485760",
        "rpc-save-upload-metadata": "true",
        "rpc-secure": "false",
        "save-not-found": "true",
        "save-session": "/home/nas-data/.aria2/aria2.session",
        "save-session-interval": "1",
        "seed-ratio": "1.0",
        "seed-time": "0",
        "server-stat-timeout": "86400",
        "show-console-readout": "false",
        "show-files": "false",
        "socket-recv-buffer-size": "0",
        "split": "64",
        "stderr": "false",
        "stop": "0",
        "stream-piece-selector": "default",
        "summary-interval": "0",
        "timeout": "10",
        "truncate-console-readout": "true",
        "uri-selector": "feedback",
        "use-head": "false",
        "user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36 Edg/93.0.961.47"
    }
	`)

	var options Options
	assert.NoError(t, json.Unmarshal(data, &options), "Couldn't unmarshal JSON")

	t.Log(options)
}
