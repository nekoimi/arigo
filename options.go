package arigo

// Options represents the aria2 input file options
type Options struct {
	AllProxy                      string  `json:"all-proxy,omitempty"`
	AllProxyPassword              string  `json:"all-proxy-passwd,omitempty"`
	AllProxyUser                  string  `json:"all-proxy-user,omitempty"`
	AllowOverwrite                bool    `json:"allow-overwrite,omitempty,string"`
	AllowPieceLengthChange        bool    `json:"allow-piece-length-change,omitempty,string"`
	AlwaysResume                  bool    `json:"always-resume,omitempty,string"`
	AsyncDNS                      bool    `json:"async-dns,omitempty,string"`
	AutoFileRenaming              bool    `json:"auto-file-renaming,omitempty,string"`
	BTEnableHookAfterHashCheck    bool    `json:"bt-enable-hook-after-hash-check,omitempty,string"`
	BTEnableLpd                   bool    `json:"bt-enable-lpd,omitempty,string"`
	BTExcludeTracker              string  `json:"bt-exclude-tracker,omitempty"`
	BTExternalIP                  string  `json:"bt-external-ip,omitempty"`
	BTForceEncryption             bool    `json:"bt-force-encryption,omitempty,string"`
	BTHashCheckSeed               bool    `json:"bt-hash-check-seed,omitempty,string"`
	BTLoadSavedMetadata           bool    `json:"bt-load-saved-metadata,omitempty,string"`
	BTMaxPeers                    uint    `json:"bt-max-peers,omitempty,string"`
	BTMetadataOnly                bool    `json:"bt-metadata-only,omitempty,string"`
	BTMinCryptoLevel              string  `json:"bt-min-crypto-level,omitempty"`
	BTPrioritizePiece             string  `json:"bt-prioritize-piece,omitempty"`
	BTRemoveUnselectedFile        bool    `json:"bt-remove-unselected-file,omitempty,string"`
	BTRequestPeerSpeedLimit       string  `json:"bt-request-peer-speed-limit,omitempty"`
	BTRequireCrypto               bool    `json:"bt-require-crypto,omitempty,string"`
	BTSaveMetadata                bool    `json:"bt-save-metadata,omitempty,string"`
	BTSeedUnverified              bool    `json:"bt-seed-unverified,omitempty,string"`
	BTStopTimeout                 uint    `json:"bt-stop-timeout,omitempty,string"`
	BTTracker                     string  `json:"bt-tracker,omitempty"`
	BTTrackerConnectTimeout       uint    `json:"bt-tracker-connect-timeout,omitempty,string"`
	BTTrackerInterval             uint    `json:"bt-tracker-interval,omitempty,string"`
	BTTrackerTimeout              uint    `json:"bt-tracker-timeout,omitempty,string"`
	CheckIntegrity                bool    `json:"check-integrity,omitempty,string"`
	Checksum                      string  `json:"checksum,omitempty"`
	ConditionalGet                bool    `json:"conditional-get,omitempty,string"`
	ConnectTimeout                uint    `json:"connect-timeout,omitempty,string"`
	ContentDispositionDefaultUTF8 bool    `json:"content-disposition-default-utf8,omitempty,string"`
	Continue                      bool    `json:"continue,omitempty,string"`
	Dir                           string  `json:"dir,omitempty"`
	DryRun                        bool    `json:"dry-run,omitempty,string"`
	EnableHTTPKeepAlive           bool    `json:"enable-http-keep-alive,omitempty,string"`
	EnableHTTPPipelining          bool    `json:"enable-http-pipelining,omitempty,string"`
	EnableMMap                    bool    `json:"enable-mmap,omitempty,string"`
	EnablePeerExchange            bool    `json:"enable-peer-exchange,omitempty,string"`
	FileAllocation                string  `json:"file-allocation,omitempty"`
	FollowMetalink                bool    `json:"follow-metalink,omitempty,string"`
	FollowTorrent                 bool    `json:"follow-torrent,omitempty,string"`
	ForceSave                     bool    `json:"force-save,omitempty,string"`
	FTPPasswd                     string  `json:"ftp-passwd,omitempty"`
	FTPPasv                       bool    `json:"ftp-pasv,omitempty,string"`
	FTPProxy                      string  `json:"ftp-proxy,omitempty"`
	FTPProxyPasswd                string  `json:"ftp-proxy-passwd,omitempty"`
	FTPProxyUser                  string  `json:"ftp-proxy-user,omitempty"`
	FTPReuseConnection            bool    `json:"ftp-reuse-connection,omitempty,string"`
	FTPType                       string  `json:"ftp-type,omitempty"`
	FTPUser                       string  `json:"ftp-user,omitempty"`
	GID                           string  `json:"gid,omitempty"`
	HashCheckOnly                 bool    `json:"hash-check-only,omitempty,string"`
	Header                        string  `json:"header,omitempty"`
	HTTPAcceptGzip                bool    `json:"http-accept-gzip,omitempty,string"`
	HTTPAuthChallenge             bool    `json:"http-auth-challenge,omitempty,string"`
	HTTPNoCache                   bool    `json:"http-no-cache,omitempty,string"`
	HTTPPasswd                    string  `json:"http-passwd,omitempty"`
	HTTPProxy                     string  `json:"http-proxy,omitempty"`
	HTTPProxyPasswd               string  `json:"http-proxy-passwd,omitempty"`
	HTTPProxyUser                 string  `json:"http-proxy-user,omitempty"`
	HTTPUser                      string  `json:"http-user,omitempty"`
	HTTPSProxy                    string  `json:"https-proxy,omitempty"`
	HTTPSProxyPasswd              string  `json:"https-proxy-passwd,omitempty"`
	HTTPSProxyUser                string  `json:"https-proxy-user,omitempty"`
	IndexOut                      uint    `json:"index-out,omitempty,string"`
	LowestSpeedLimit              string  `json:"lowest-speed-limit,omitempty"`
	MaxConcurrentDownloads        uint    `json:"max-concurrent-downloads,omitempty,string"`
	MaxConnectionPerServer        uint    `json:"max-connection-per-server,omitempty,string"`
	MaxDownloadLimit              string  `json:"max-download-limit,omitempty"`
	MaxFileNotFound               uint    `json:"max-file-not-found,omitempty,string"`
	MaxMMapLimit                  string  `json:"max-mmap-limit,omitempty"`
	MaxResumeFailureTries         uint    `json:"max-resume-failure-tries,omitempty,string"`
	MaxTries                      uint    `json:"max-tries,omitempty,string"`
	MaxUploadLimit                string  `json:"max-upload-limit,omitempty"`
	MetalinkBaseURI               string  `json:"metalink-base-uri,omitempty"`
	MetalinkEnableUniqueProtocol  bool    `json:"metalink-enable-unique-protocol,omitempty,string"`
	MetalinkLanguage              string  `json:"metalink-language,omitempty"`
	MetalinkLocation              string  `json:"metalink-location,omitempty"`
	MetalinkOS                    string  `json:"metalink-os,omitempty"`
	MetalinkPreferredProtocol     string  `json:"metalink-preferred-protocol,omitempty"`
	MetalinkVersion               string  `json:"metalink-version,omitempty"`
	MinSplitSize                  string  `json:"min-split-size,omitempty"`
	NoFileAllocationLimit         string  `json:"no-file-allocation-limit,omitempty"`
	NoNetrc                       bool    `json:"no-netrc,omitempty,string"`
	NoProxy                       bool    `json:"no-proxy,omitempty,string"`
	Out                           string  `json:"out,omitempty"`
	ParameterizedURI              string  `json:"parameterized-uri,omitempty"`
	Pause                         bool    `json:"pause,omitempty,string"`
	PauseMetadata                 bool    `json:"pause-metadata,omitempty,string"`
	PieceLength                   string  `json:"piece-length,omitempty"`
	ProxyMethod                   string  `json:"proxy-method,omitempty"`
	RealtimeChunkChecksum         string  `json:"realtime-chunk-checksum,omitempty"`
	Referer                       string  `json:"referer,omitempty"`
	RemoteTime                    bool    `json:"remote-time,omitempty,string"`
	RemoveControlFile             string  `json:"remove-control-file,omitempty"`
	RetryWait                     uint    `json:"retry-wait,omitempty,string"`
	ReuseURI                      bool    `json:"reuse-uri,omitempty,string"`
	RPCSaveUploadMetadata         string  `json:"rpc-save-upload-metadata,omitempty"`
	SeedRatio                     float32 `json:"seed-ratio,omitempty,string"`
	SeedTime                      uint    `json:"seed-time,omitempty,string"`
	SelectFile                    string  `json:"select-file,omitempty"`
	Split                         uint    `json:"split,omitempty,string"`
	SSHHostKeyMD                  string  `json:"ssh-host-key-md,omitempty"`
	StreamPieceSelector           string  `json:"stream-piece-selector,omitempty"`
	Timeout                       uint    `json:"timeout,omitempty,string"`
	URISelector                   string  `json:"uri-selector,omitempty"`
	UseHead                       bool    `json:"use-head,omitempty,string"`
	UserAgent                     string  `json:"user-agent,omitempty"`
}
