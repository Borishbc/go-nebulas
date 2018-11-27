// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config.proto

package nebletpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Reporting modules.
type StatsConfig_ReportingModule int32

const (
	StatsConfig_Influxdb StatsConfig_ReportingModule = 0
)

var StatsConfig_ReportingModule_name = map[int32]string{
	0: "Influxdb",
}
var StatsConfig_ReportingModule_value = map[string]int32{
	"Influxdb": 0,
}

func (x StatsConfig_ReportingModule) String() string {
	return proto.EnumName(StatsConfig_ReportingModule_name, int32(x))
}
func (StatsConfig_ReportingModule) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_config_31b8b1003b247562, []int{7, 0}
}

// Neblet global configurations.
type Config struct {
	// Network config.
	Network *NetworkConfig `protobuf:"bytes,1,opt,name=network" json:"network"`
	// Chain config.
	Chain *ChainConfig `protobuf:"bytes,2,opt,name=chain" json:"chain"`
	// RPC config.
	Rpc *RPCConfig `protobuf:"bytes,3,opt,name=rpc" json:"rpc"`
	// Stats config.
	Stats *StatsConfig `protobuf:"bytes,100,opt,name=stats" json:"stats"`
	// Misc config.
	Misc *MiscConfig `protobuf:"bytes,101,opt,name=misc" json:"misc"`
	// App Config.
	App                  *AppConfig `protobuf:"bytes,102,opt,name=app" json:"app"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_31b8b1003b247562, []int{0}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (dst *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(dst, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetNetwork() *NetworkConfig {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *Config) GetChain() *ChainConfig {
	if m != nil {
		return m.Chain
	}
	return nil
}

func (m *Config) GetRpc() *RPCConfig {
	if m != nil {
		return m.Rpc
	}
	return nil
}

func (m *Config) GetStats() *StatsConfig {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *Config) GetMisc() *MiscConfig {
	if m != nil {
		return m.Misc
	}
	return nil
}

func (m *Config) GetApp() *AppConfig {
	if m != nil {
		return m.App
	}
	return nil
}

type NetworkConfig struct {
	// Neb seed node address.
	Seed []string `protobuf:"bytes,1,rep,name=seed" json:"seed"`
	// Listen addresses.
	Listen []string `protobuf:"bytes,2,rep,name=listen" json:"listen"`
	// Network node privateKey address. If nil, generate a new node.
	PrivateKey string `protobuf:"bytes,3,opt,name=private_key,json=privateKey,proto3" json:"private_key"`
	// Network ID
	NetworkId            uint32   `protobuf:"varint,4,opt,name=network_id,json=networkId,proto3" json:"network_id"`
	StreamLimits         int32    `protobuf:"varint,5,opt,name=stream_limits,json=streamLimits,proto3" json:"stream_limits"`
	ReservedStreamLimits int32    `protobuf:"varint,6,opt,name=reserved_stream_limits,json=reservedStreamLimits,proto3" json:"reserved_stream_limits"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkConfig) Reset()         { *m = NetworkConfig{} }
func (m *NetworkConfig) String() string { return proto.CompactTextString(m) }
func (*NetworkConfig) ProtoMessage()    {}
func (*NetworkConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_31b8b1003b247562, []int{1}
}
func (m *NetworkConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkConfig.Unmarshal(m, b)
}
func (m *NetworkConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkConfig.Marshal(b, m, deterministic)
}
func (dst *NetworkConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkConfig.Merge(dst, src)
}
func (m *NetworkConfig) XXX_Size() int {
	return xxx_messageInfo_NetworkConfig.Size(m)
}
func (m *NetworkConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkConfig proto.InternalMessageInfo

func (m *NetworkConfig) GetSeed() []string {
	if m != nil {
		return m.Seed
	}
	return nil
}

func (m *NetworkConfig) GetListen() []string {
	if m != nil {
		return m.Listen
	}
	return nil
}

func (m *NetworkConfig) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *NetworkConfig) GetNetworkId() uint32 {
	if m != nil {
		return m.NetworkId
	}
	return 0
}

func (m *NetworkConfig) GetStreamLimits() int32 {
	if m != nil {
		return m.StreamLimits
	}
	return 0
}

func (m *NetworkConfig) GetReservedStreamLimits() int32 {
	if m != nil {
		return m.ReservedStreamLimits
	}
	return 0
}

type ChainConfig struct {
	// ChainID.
	ChainId uint32 `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id"`
	// genesis conf file path
	Genesis string `protobuf:"bytes,2,opt,name=genesis,proto3" json:"genesis"`
	// Data dir.
	Datadir string `protobuf:"bytes,11,opt,name=datadir,proto3" json:"datadir"`
	// Key dir.
	Keydir string `protobuf:"bytes,12,opt,name=keydir,proto3" json:"keydir"`
	// Start mine at launch
	StartMine bool `protobuf:"varint,20,opt,name=start_mine,json=startMine,proto3" json:"start_mine"`
	// Coinbase.
	Coinbase string `protobuf:"bytes,21,opt,name=coinbase,proto3" json:"coinbase"`
	// Miner.
	Miner string `protobuf:"bytes,22,opt,name=miner,proto3" json:"miner"`
	// Passphrase.
	Passphrase string `protobuf:"bytes,23,opt,name=passphrase,proto3" json:"passphrase"`
	// Enable remote sign server
	EnableRemoteSignServer bool `protobuf:"varint,24,opt,name=enable_remote_sign_server,json=enableRemoteSignServer,proto3" json:"enable_remote_sign_server"`
	// Remote sign server
	RemoteSignServer string `protobuf:"bytes,25,opt,name=remote_sign_server,json=remoteSignServer,proto3" json:"remote_sign_server"`
	// Lowest GasPrice.
	GasPrice string `protobuf:"bytes,26,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price"`
	// Max GasLimit.
	GasLimit string `protobuf:"bytes,27,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit"`
	// Supported signature cipher list. ["ECC_SECP256K1"]
	SignatureCiphers   []string `protobuf:"bytes,28,rep,name=signature_ciphers,json=signatureCiphers" json:"signature_ciphers"`
	SuperNode          bool     `protobuf:"varint,30,opt,name=super_node,json=superNode,proto3" json:"super_node"`
	UnsupportedKeyword string   `protobuf:"bytes,31,opt,name=unsupported_keyword,json=unsupportedKeyword,proto3" json:"unsupported_keyword"`
	Dynasty            string   `protobuf:"bytes,32,opt,name=dynasty,proto3" json:"dynasty"`
	// access control config path
	Access string `protobuf:"bytes,33,opt,name=access,proto3" json:"access"`
	// v8 server path
	NvmPath string `protobuf:"bytes,34,opt,name=nvm_path,json=nvmPath,proto3" json:"nvm_path"`
	// v8 server listen addresses
	NvmListen            []string `protobuf:"bytes,35,rep,name=nvm_listen,json=nvmListen" json:"nvm_listen"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChainConfig) Reset()         { *m = ChainConfig{} }
func (m *ChainConfig) String() string { return proto.CompactTextString(m) }
func (*ChainConfig) ProtoMessage()    {}
func (*ChainConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_31b8b1003b247562, []int{2}
}
func (m *ChainConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainConfig.Unmarshal(m, b)
}
func (m *ChainConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainConfig.Marshal(b, m, deterministic)
}
func (dst *ChainConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainConfig.Merge(dst, src)
}
func (m *ChainConfig) XXX_Size() int {
	return xxx_messageInfo_ChainConfig.Size(m)
}
func (m *ChainConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ChainConfig proto.InternalMessageInfo

func (m *ChainConfig) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *ChainConfig) GetGenesis() string {
	if m != nil {
		return m.Genesis
	}
	return ""
}

func (m *ChainConfig) GetDatadir() string {
	if m != nil {
		return m.Datadir
	}
	return ""
}

func (m *ChainConfig) GetKeydir() string {
	if m != nil {
		return m.Keydir
	}
	return ""
}

func (m *ChainConfig) GetStartMine() bool {
	if m != nil {
		return m.StartMine
	}
	return false
}

func (m *ChainConfig) GetCoinbase() string {
	if m != nil {
		return m.Coinbase
	}
	return ""
}

func (m *ChainConfig) GetMiner() string {
	if m != nil {
		return m.Miner
	}
	return ""
}

func (m *ChainConfig) GetPassphrase() string {
	if m != nil {
		return m.Passphrase
	}
	return ""
}

func (m *ChainConfig) GetEnableRemoteSignServer() bool {
	if m != nil {
		return m.EnableRemoteSignServer
	}
	return false
}

func (m *ChainConfig) GetRemoteSignServer() string {
	if m != nil {
		return m.RemoteSignServer
	}
	return ""
}

func (m *ChainConfig) GetGasPrice() string {
	if m != nil {
		return m.GasPrice
	}
	return ""
}

func (m *ChainConfig) GetGasLimit() string {
	if m != nil {
		return m.GasLimit
	}
	return ""
}

func (m *ChainConfig) GetSignatureCiphers() []string {
	if m != nil {
		return m.SignatureCiphers
	}
	return nil
}

func (m *ChainConfig) GetSuperNode() bool {
	if m != nil {
		return m.SuperNode
	}
	return false
}

func (m *ChainConfig) GetUnsupportedKeyword() string {
	if m != nil {
		return m.UnsupportedKeyword
	}
	return ""
}

func (m *ChainConfig) GetDynasty() string {
	if m != nil {
		return m.Dynasty
	}
	return ""
}

func (m *ChainConfig) GetAccess() string {
	if m != nil {
		return m.Access
	}
	return ""
}

func (m *ChainConfig) GetNvmPath() string {
	if m != nil {
		return m.NvmPath
	}
	return ""
}

func (m *ChainConfig) GetNvmListen() []string {
	if m != nil {
		return m.NvmListen
	}
	return nil
}

type RPCConfig struct {
	// RPC listen addresses.
	RpcListen []string `protobuf:"bytes,1,rep,name=rpc_listen,json=rpcListen" json:"rpc_listen"`
	// HTTP listen addresses.
	HttpListen []string `protobuf:"bytes,2,rep,name=http_listen,json=httpListen" json:"http_listen"`
	// Enabled HTTP modules.["api", "admin"]
	HttpModule       []string `protobuf:"bytes,3,rep,name=http_module,json=httpModule" json:"http_module"`
	ConnectionLimits int32    `protobuf:"varint,4,opt,name=connection_limits,json=connectionLimits,proto3" json:"connection_limits"`
	HttpLimits       int32    `protobuf:"varint,5,opt,name=http_limits,json=httpLimits,proto3" json:"http_limits"`
	// HTTP CORS allowed origins
	HttpCors             []string `protobuf:"bytes,6,rep,name=http_cors,json=httpCors" json:"http_cors"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RPCConfig) Reset()         { *m = RPCConfig{} }
func (m *RPCConfig) String() string { return proto.CompactTextString(m) }
func (*RPCConfig) ProtoMessage()    {}
func (*RPCConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_31b8b1003b247562, []int{3}
}
func (m *RPCConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCConfig.Unmarshal(m, b)
}
func (m *RPCConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCConfig.Marshal(b, m, deterministic)
}
func (dst *RPCConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCConfig.Merge(dst, src)
}
func (m *RPCConfig) XXX_Size() int {
	return xxx_messageInfo_RPCConfig.Size(m)
}
func (m *RPCConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RPCConfig proto.InternalMessageInfo

func (m *RPCConfig) GetRpcListen() []string {
	if m != nil {
		return m.RpcListen
	}
	return nil
}

func (m *RPCConfig) GetHttpListen() []string {
	if m != nil {
		return m.HttpListen
	}
	return nil
}

func (m *RPCConfig) GetHttpModule() []string {
	if m != nil {
		return m.HttpModule
	}
	return nil
}

func (m *RPCConfig) GetConnectionLimits() int32 {
	if m != nil {
		return m.ConnectionLimits
	}
	return 0
}

func (m *RPCConfig) GetHttpLimits() int32 {
	if m != nil {
		return m.HttpLimits
	}
	return 0
}

func (m *RPCConfig) GetHttpCors() []string {
	if m != nil {
		return m.HttpCors
	}
	return nil
}

type AppConfig struct {
	LogLevel string `protobuf:"bytes,1,opt,name=log_level,json=logLevel,proto3" json:"log_level"`
	LogFile  string `protobuf:"bytes,2,opt,name=log_file,json=logFile,proto3" json:"log_file"`
	// log file age, unit is s.
	LogAge            uint32 `protobuf:"varint,3,opt,name=log_age,json=logAge,proto3" json:"log_age"`
	EnableCrashReport bool   `protobuf:"varint,4,opt,name=enable_crash_report,json=enableCrashReport,proto3" json:"enable_crash_report"`
	CrashReportUrl    string `protobuf:"bytes,5,opt,name=crash_report_url,json=crashReportUrl,proto3" json:"crash_report_url"`
	// pprof config
	Pprof                *PprofConfig `protobuf:"bytes,6,opt,name=pprof" json:"pprof"`
	Version              string       `protobuf:"bytes,100,opt,name=version,proto3" json:"version"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AppConfig) Reset()         { *m = AppConfig{} }
func (m *AppConfig) String() string { return proto.CompactTextString(m) }
func (*AppConfig) ProtoMessage()    {}
func (*AppConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_31b8b1003b247562, []int{4}
}
func (m *AppConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppConfig.Unmarshal(m, b)
}
func (m *AppConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppConfig.Marshal(b, m, deterministic)
}
func (dst *AppConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppConfig.Merge(dst, src)
}
func (m *AppConfig) XXX_Size() int {
	return xxx_messageInfo_AppConfig.Size(m)
}
func (m *AppConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AppConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AppConfig proto.InternalMessageInfo

func (m *AppConfig) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

func (m *AppConfig) GetLogFile() string {
	if m != nil {
		return m.LogFile
	}
	return ""
}

func (m *AppConfig) GetLogAge() uint32 {
	if m != nil {
		return m.LogAge
	}
	return 0
}

func (m *AppConfig) GetEnableCrashReport() bool {
	if m != nil {
		return m.EnableCrashReport
	}
	return false
}

func (m *AppConfig) GetCrashReportUrl() string {
	if m != nil {
		return m.CrashReportUrl
	}
	return ""
}

func (m *AppConfig) GetPprof() *PprofConfig {
	if m != nil {
		return m.Pprof
	}
	return nil
}

func (m *AppConfig) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type PprofConfig struct {
	// pprof listen address, if not configured, the function closes.
	HttpListen string `protobuf:"bytes,1,opt,name=http_listen,json=httpListen,proto3" json:"http_listen"`
	// cpu profiling file, if not configured, the profiling not start
	Cpuprofile string `protobuf:"bytes,2,opt,name=cpuprofile,proto3" json:"cpuprofile"`
	// memory profiling file, if not configured, the profiling not start
	Memprofile           string   `protobuf:"bytes,3,opt,name=memprofile,proto3" json:"memprofile"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PprofConfig) Reset()         { *m = PprofConfig{} }
func (m *PprofConfig) String() string { return proto.CompactTextString(m) }
func (*PprofConfig) ProtoMessage()    {}
func (*PprofConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_31b8b1003b247562, []int{5}
}
func (m *PprofConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PprofConfig.Unmarshal(m, b)
}
func (m *PprofConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PprofConfig.Marshal(b, m, deterministic)
}
func (dst *PprofConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PprofConfig.Merge(dst, src)
}
func (m *PprofConfig) XXX_Size() int {
	return xxx_messageInfo_PprofConfig.Size(m)
}
func (m *PprofConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PprofConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PprofConfig proto.InternalMessageInfo

func (m *PprofConfig) GetHttpListen() string {
	if m != nil {
		return m.HttpListen
	}
	return ""
}

func (m *PprofConfig) GetCpuprofile() string {
	if m != nil {
		return m.Cpuprofile
	}
	return ""
}

func (m *PprofConfig) GetMemprofile() string {
	if m != nil {
		return m.Memprofile
	}
	return ""
}

type MiscConfig struct {
	// Default encryption ciper when create new keystore file.
	DefaultKeystoreFileCiper string   `protobuf:"bytes,1,opt,name=default_keystore_file_ciper,json=defaultKeystoreFileCiper,proto3" json:"default_keystore_file_ciper"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *MiscConfig) Reset()         { *m = MiscConfig{} }
func (m *MiscConfig) String() string { return proto.CompactTextString(m) }
func (*MiscConfig) ProtoMessage()    {}
func (*MiscConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_31b8b1003b247562, []int{6}
}
func (m *MiscConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MiscConfig.Unmarshal(m, b)
}
func (m *MiscConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MiscConfig.Marshal(b, m, deterministic)
}
func (dst *MiscConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MiscConfig.Merge(dst, src)
}
func (m *MiscConfig) XXX_Size() int {
	return xxx_messageInfo_MiscConfig.Size(m)
}
func (m *MiscConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MiscConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MiscConfig proto.InternalMessageInfo

func (m *MiscConfig) GetDefaultKeystoreFileCiper() string {
	if m != nil {
		return m.DefaultKeystoreFileCiper
	}
	return ""
}

type StatsConfig struct {
	// Enable metrics or not.
	EnableMetrics   bool                          `protobuf:"varint,1,opt,name=enable_metrics,json=enableMetrics,proto3" json:"enable_metrics"`
	ReportingModule []StatsConfig_ReportingModule `protobuf:"varint,2,rep,packed,name=reporting_module,json=reportingModule,enum=nebletpb.StatsConfig_ReportingModule" json:"reporting_module"`
	// Influxdb config.
	Influxdb             *InfluxdbConfig `protobuf:"bytes,11,opt,name=influxdb" json:"influxdb"`
	MetricsTags          []string        `protobuf:"bytes,12,rep,name=metrics_tags,json=metricsTags" json:"metrics_tags"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StatsConfig) Reset()         { *m = StatsConfig{} }
func (m *StatsConfig) String() string { return proto.CompactTextString(m) }
func (*StatsConfig) ProtoMessage()    {}
func (*StatsConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_31b8b1003b247562, []int{7}
}
func (m *StatsConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsConfig.Unmarshal(m, b)
}
func (m *StatsConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsConfig.Marshal(b, m, deterministic)
}
func (dst *StatsConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsConfig.Merge(dst, src)
}
func (m *StatsConfig) XXX_Size() int {
	return xxx_messageInfo_StatsConfig.Size(m)
}
func (m *StatsConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsConfig.DiscardUnknown(m)
}

var xxx_messageInfo_StatsConfig proto.InternalMessageInfo

func (m *StatsConfig) GetEnableMetrics() bool {
	if m != nil {
		return m.EnableMetrics
	}
	return false
}

func (m *StatsConfig) GetReportingModule() []StatsConfig_ReportingModule {
	if m != nil {
		return m.ReportingModule
	}
	return nil
}

func (m *StatsConfig) GetInfluxdb() *InfluxdbConfig {
	if m != nil {
		return m.Influxdb
	}
	return nil
}

func (m *StatsConfig) GetMetricsTags() []string {
	if m != nil {
		return m.MetricsTags
	}
	return nil
}

type InfluxdbConfig struct {
	// Host.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host"`
	// Port.
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port"`
	// Database name.
	Db string `protobuf:"bytes,3,opt,name=db,proto3" json:"db"`
	// Auth user.
	User string `protobuf:"bytes,4,opt,name=user,proto3" json:"user"`
	// Auth password.
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfluxdbConfig) Reset()         { *m = InfluxdbConfig{} }
func (m *InfluxdbConfig) String() string { return proto.CompactTextString(m) }
func (*InfluxdbConfig) ProtoMessage()    {}
func (*InfluxdbConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_31b8b1003b247562, []int{8}
}
func (m *InfluxdbConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfluxdbConfig.Unmarshal(m, b)
}
func (m *InfluxdbConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfluxdbConfig.Marshal(b, m, deterministic)
}
func (dst *InfluxdbConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfluxdbConfig.Merge(dst, src)
}
func (m *InfluxdbConfig) XXX_Size() int {
	return xxx_messageInfo_InfluxdbConfig.Size(m)
}
func (m *InfluxdbConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_InfluxdbConfig.DiscardUnknown(m)
}

var xxx_messageInfo_InfluxdbConfig proto.InternalMessageInfo

func (m *InfluxdbConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *InfluxdbConfig) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *InfluxdbConfig) GetDb() string {
	if m != nil {
		return m.Db
	}
	return ""
}

func (m *InfluxdbConfig) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *InfluxdbConfig) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "nebletpb.Config")
	proto.RegisterType((*NetworkConfig)(nil), "nebletpb.NetworkConfig")
	proto.RegisterType((*ChainConfig)(nil), "nebletpb.ChainConfig")
	proto.RegisterType((*RPCConfig)(nil), "nebletpb.RPCConfig")
	proto.RegisterType((*AppConfig)(nil), "nebletpb.AppConfig")
	proto.RegisterType((*PprofConfig)(nil), "nebletpb.PprofConfig")
	proto.RegisterType((*MiscConfig)(nil), "nebletpb.MiscConfig")
	proto.RegisterType((*StatsConfig)(nil), "nebletpb.StatsConfig")
	proto.RegisterType((*InfluxdbConfig)(nil), "nebletpb.InfluxdbConfig")
	proto.RegisterEnum("nebletpb.StatsConfig_ReportingModule", StatsConfig_ReportingModule_name, StatsConfig_ReportingModule_value)
}

func init() { proto.RegisterFile("config.proto", fileDescriptor_config_31b8b1003b247562) }

var fileDescriptor_config_31b8b1003b247562 = []byte{
	// 1058 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x56, 0xdb, 0x6e, 0x1b, 0x37,
	0x13, 0xfe, 0xe5, 0x53, 0xb4, 0x23, 0xdb, 0x71, 0x18, 0xc7, 0x61, 0xe2, 0xbf, 0x89, 0xa3, 0x20,
	0x80, 0x80, 0x14, 0x2e, 0x9a, 0xe6, 0xa6, 0x17, 0xbd, 0x08, 0x04, 0x14, 0x08, 0x6c, 0x07, 0xc6,
	0xba, 0xbd, 0x5e, 0x50, 0xbb, 0xa3, 0x15, 0x91, 0x5d, 0x2e, 0x41, 0x52, 0x4a, 0x7c, 0xd7, 0x17,
	0xe8, 0xeb, 0xb5, 0xcf, 0xd1, 0x07, 0x28, 0x50, 0xcc, 0x2c, 0x57, 0x27, 0xe4, 0x8e, 0xf3, 0x7d,
	0xdf, 0x70, 0x96, 0x73, 0x92, 0xe0, 0x30, 0x6f, 0xcc, 0x54, 0x97, 0x97, 0xd6, 0x35, 0xa1, 0x11,
	0x7d, 0x83, 0x93, 0x0a, 0x83, 0x9d, 0x0c, 0xff, 0xdc, 0x81, 0x83, 0x31, 0x53, 0xe2, 0x47, 0x78,
	0x60, 0x30, 0x7c, 0x69, 0xdc, 0x67, 0xd9, 0xbb, 0xe8, 0x8d, 0x06, 0xef, 0x9e, 0x5e, 0x76, 0xb2,
	0xcb, 0x4f, 0x2d, 0xd1, 0x2a, 0xd3, 0x4e, 0x27, 0xde, 0xc2, 0x7e, 0x3e, 0x53, 0xda, 0xc8, 0x1d,
	0x76, 0x78, 0xb2, 0x72, 0x18, 0x13, 0x1c, 0xe5, 0xad, 0x46, 0xbc, 0x81, 0x5d, 0x67, 0x73, 0xb9,
	0xcb, 0xd2, 0xc7, 0x2b, 0x69, 0x7a, 0x3b, 0x8e, 0x42, 0xe2, 0xe9, 0x4e, 0x1f, 0x54, 0xf0, 0xb2,
	0xd8, 0xbe, 0xf3, 0x8e, 0xe0, 0xee, 0x4e, 0xd6, 0x88, 0x11, 0xec, 0xd5, 0xda, 0xe7, 0x12, 0x59,
	0x7b, 0xba, 0xd2, 0xde, 0x68, 0x9f, 0x47, 0x29, 0x2b, 0x28, 0xba, 0xb2, 0x56, 0x4e, 0xb7, 0xa3,
	0x7f, 0xb0, 0xb6, 0x8b, 0xae, 0xac, 0x1d, 0xfe, 0xd5, 0x83, 0xa3, 0x8d, 0xc7, 0x0a, 0x01, 0x7b,
	0x1e, 0xb1, 0x90, 0xbd, 0x8b, 0xdd, 0x51, 0x92, 0xf2, 0x59, 0x9c, 0xc1, 0x41, 0xa5, 0x7d, 0x40,
	0x7a, 0x38, 0xa1, 0xd1, 0x12, 0x2f, 0x61, 0x60, 0x9d, 0x5e, 0xa8, 0x80, 0xd9, 0x67, 0xbc, 0xe7,
	0xa7, 0x26, 0x29, 0x44, 0xe8, 0x0a, 0xef, 0xc5, 0x77, 0x00, 0x31, 0x77, 0x99, 0x2e, 0xe4, 0xde,
	0x45, 0x6f, 0x74, 0x94, 0x26, 0x11, 0xf9, 0x58, 0x88, 0xd7, 0x70, 0xe4, 0x83, 0x43, 0x55, 0x67,
	0x95, 0xae, 0x75, 0xf0, 0x72, 0xff, 0xa2, 0x37, 0xda, 0x4f, 0x0f, 0x5b, 0xf0, 0x9a, 0x31, 0xf1,
	0x1e, 0xce, 0x1c, 0x7a, 0x74, 0x0b, 0x2c, 0xb2, 0x4d, 0xf5, 0x01, 0xab, 0x4f, 0x3b, 0xf6, 0x6e,
	0xcd, 0x6b, 0xf8, 0xcf, 0x1e, 0x0c, 0xd6, 0x8a, 0x22, 0x9e, 0x41, 0x9f, 0xcb, 0x42, 0xdf, 0xd1,
	0xe3, 0xef, 0x78, 0xc0, 0xf6, 0xc7, 0x42, 0x48, 0x78, 0x50, 0xa2, 0x41, 0xaf, 0x3d, 0xd7, 0x35,
	0x49, 0x3b, 0x93, 0x98, 0x42, 0x05, 0x55, 0x68, 0x27, 0x07, 0x2d, 0x13, 0x4d, 0xca, 0xc8, 0x67,
	0xbc, 0x27, 0xe2, 0x90, 0x89, 0x68, 0xd1, 0x83, 0x7d, 0x50, 0x2e, 0x64, 0xb5, 0x36, 0x28, 0x4f,
	0x2f, 0x7a, 0xa3, 0x7e, 0x9a, 0x30, 0x72, 0xa3, 0x0d, 0x8a, 0xe7, 0xd0, 0xcf, 0x1b, 0x6d, 0x26,
	0xca, 0xa3, 0x7c, 0xc2, 0x8e, 0x4b, 0x5b, 0x9c, 0xc2, 0x3e, 0x39, 0x39, 0x79, 0xc6, 0x44, 0x6b,
	0x88, 0x17, 0x00, 0x56, 0x79, 0x6f, 0x67, 0x8e, 0x7c, 0x9e, 0xc6, 0x0c, 0x2f, 0x11, 0xf1, 0x33,
	0x3c, 0x43, 0xa3, 0x26, 0x15, 0x66, 0x0e, 0xeb, 0x26, 0x60, 0xe6, 0x75, 0x69, 0x32, 0x4e, 0x88,
	0x93, 0x92, 0xe3, 0x9f, 0xb5, 0x82, 0x94, 0xf9, 0x3b, 0x5d, 0x9a, 0x3b, 0x66, 0xc5, 0xf7, 0x20,
	0xbe, 0xe1, 0xf3, 0x8c, 0x43, 0x9c, 0xb8, 0x6d, 0xf5, 0x39, 0x24, 0xa5, 0xf2, 0x99, 0x75, 0x3a,
	0x47, 0xf9, 0xbc, 0xfd, 0xf6, 0x52, 0xf9, 0x5b, 0xb2, 0x3b, 0x92, 0xeb, 0x22, 0xcf, 0x97, 0x24,
	0xd7, 0x42, 0xbc, 0x85, 0x47, 0x14, 0x40, 0x85, 0xb9, 0xc3, 0x2c, 0xd7, 0x76, 0x86, 0xce, 0xcb,
	0xff, 0x73, 0x23, 0x9d, 0x2c, 0x89, 0x71, 0x8b, 0x73, 0x02, 0xe7, 0x16, 0x5d, 0x66, 0x9a, 0x02,
	0xe5, 0x8b, 0x98, 0x40, 0x42, 0x3e, 0x35, 0x05, 0x8a, 0x1f, 0xe0, 0xf1, 0xdc, 0xf8, 0xb9, 0xb5,
	0x8d, 0x0b, 0x58, 0x50, 0xd7, 0x7d, 0x69, 0x5c, 0x21, 0x5f, 0x72, 0x48, 0xb1, 0x46, 0x5d, 0xb5,
	0x0c, 0x97, 0xf0, 0xde, 0x28, 0x1f, 0xee, 0xe5, 0x45, 0x2c, 0x61, 0x6b, 0x52, 0x09, 0x55, 0x9e,
	0xa3, 0xf7, 0xf2, 0x55, 0x5b, 0xc2, 0xd6, 0xa2, 0x4e, 0x31, 0x8b, 0x3a, 0xb3, 0x2a, 0xcc, 0xe4,
	0xb0, 0x75, 0x31, 0x8b, 0xfa, 0x56, 0x85, 0x19, 0xb7, 0xf3, 0x82, 0xda, 0x8f, 0x67, 0xe1, 0x35,
	0x3f, 0x21, 0x31, 0x8b, 0xfa, 0x9a, 0x81, 0xe1, 0xdf, 0x3d, 0x48, 0x96, 0xd3, 0x4d, 0x62, 0x67,
	0xf3, 0x4e, 0xdc, 0x8e, 0x53, 0xe2, 0x6c, 0x7e, 0xbd, 0x9c, 0x9d, 0x59, 0x08, 0x36, 0xdb, 0x18,
	0x2c, 0x20, 0x68, 0x4b, 0x50, 0x37, 0xc5, 0xbc, 0x42, 0xb9, 0xbb, 0x12, 0xdc, 0x30, 0x42, 0x79,
	0xcd, 0x1b, 0x63, 0x30, 0x0f, 0xba, 0x31, 0xdd, 0x4c, 0xec, 0xf1, 0x4c, 0x9c, 0xac, 0x88, 0x38,
	0x45, 0xab, 0x70, 0x6b, 0x83, 0x16, 0xc3, 0xb1, 0xe0, 0x1c, 0x12, 0x16, 0xe4, 0x8d, 0xa3, 0xc9,
	0xa2, 0x60, 0x7d, 0x02, 0xc6, 0x8d, 0xf3, 0xc3, 0x7f, 0x7b, 0x90, 0x2c, 0x37, 0x07, 0x49, 0xab,
	0xa6, 0xcc, 0x2a, 0x5c, 0x60, 0xc5, 0xc3, 0x94, 0xa4, 0xfd, 0xaa, 0x29, 0xaf, 0xc9, 0xa6, 0xf4,
	0x11, 0x39, 0xd5, 0x15, 0x76, 0xe3, 0x54, 0x35, 0xe5, 0xaf, 0xba, 0x42, 0xf1, 0x14, 0xe8, 0x98,
	0xa9, 0x12, 0x79, 0x55, 0x1c, 0xa5, 0x07, 0x55, 0x53, 0x7e, 0x28, 0x51, 0x5c, 0xc2, 0xe3, 0xd8,
	0xc4, 0xb9, 0x53, 0x7e, 0x96, 0x39, 0xa4, 0x22, 0xf2, 0x5b, 0xfa, 0xe9, 0xa3, 0x96, 0x1a, 0x13,
	0x93, 0x32, 0x21, 0x46, 0x70, 0xb2, 0x2e, 0xcc, 0xe6, 0xae, 0xe2, 0x17, 0x25, 0xe9, 0x71, 0xbe,
	0x92, 0xfd, 0xee, 0x2a, 0xda, 0xae, 0xd6, 0xba, 0x66, 0xca, 0xbb, 0x62, 0x63, 0xbb, 0xde, 0x12,
	0xdc, 0x6d, 0x57, 0xd6, 0x50, 0xaf, 0x2c, 0xd0, 0x79, 0xdd, 0x18, 0x5e, 0xc6, 0x49, 0xda, 0x99,
	0x43, 0x03, 0x83, 0x35, 0xfd, 0x76, 0xed, 0xda, 0x14, 0xac, 0xd7, 0xee, 0x05, 0x40, 0x6e, 0xe7,
	0xe4, 0xb1, 0x4a, 0xc3, 0x1a, 0x42, 0x7c, 0x8d, 0x75, 0xc7, 0xc7, 0xbd, 0xb9, 0x42, 0x86, 0x57,
	0x00, 0xab, 0x8d, 0x2e, 0x7e, 0x81, 0xf3, 0x02, 0xa7, 0x6a, 0x5e, 0x05, 0x6a, 0x78, 0x1f, 0x1a,
	0x87, 0x9c, 0x5f, 0x1a, 0x26, 0x74, 0x31, 0xbc, 0x8c, 0x92, 0xab, 0xa8, 0xa0, 0x8c, 0x8f, 0x89,
	0x1f, 0xfe, 0xb1, 0x03, 0x83, 0xb5, 0xdf, 0x12, 0xf1, 0x06, 0x8e, 0x63, 0xb6, 0x6b, 0x0c, 0x4e,
	0xe7, 0x9e, 0x6f, 0xe8, 0xa7, 0x47, 0x2d, 0x7a, 0xd3, 0x82, 0xe2, 0x16, 0x4e, 0xda, 0xf4, 0x6a,
	0x53, 0x76, 0x4d, 0x48, 0x5d, 0x7a, 0xfc, 0xee, 0xcd, 0x37, 0x7f, 0xa3, 0x2e, 0xd3, 0x4e, 0xdd,
	0xf6, 0x67, 0xfa, 0xd0, 0x6d, 0x02, 0xe2, 0x3d, 0xf4, 0xb5, 0x99, 0x56, 0xf3, 0xaf, 0xc5, 0x84,
	0xf7, 0xe9, 0xe0, 0x9d, 0x5c, 0xdd, 0xf4, 0x31, 0x32, 0xb1, 0x24, 0x4b, 0xa5, 0x78, 0x05, 0x87,
	0xf1, 0x3b, 0xb3, 0xa0, 0x4a, 0x2f, 0x0f, 0xb9, 0x37, 0x07, 0x11, 0xfb, 0x4d, 0x95, 0x7e, 0xf8,
	0x12, 0x1e, 0x6e, 0x05, 0x17, 0x87, 0xd0, 0xef, 0x6e, 0x3c, 0xf9, 0xdf, 0xf0, 0x2b, 0x1c, 0x6f,
	0xde, 0x4f, 0x3f, 0x73, 0xb3, 0xc6, 0x87, 0x98, 0x3c, 0x3e, 0x13, 0xc6, 0x7d, 0xb7, 0xc3, 0xcd,
	0xc9, 0x67, 0x71, 0x0c, 0x3b, 0xc5, 0x24, 0x56, 0x68, 0xa7, 0x98, 0x90, 0x66, 0xee, 0xd1, 0x71,
	0x6f, 0x26, 0x29, 0x9f, 0x69, 0xab, 0xd3, 0x46, 0xe6, 0x4d, 0xd4, 0xb6, 0xe1, 0xd2, 0x9e, 0x1c,
	0xf0, 0x3f, 0x90, 0x9f, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xe6, 0x63, 0x58, 0x91, 0x08,
	0x00, 0x00,
}
