// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config.proto

/*
Package nebletpb is a generated protocol buffer package.

It is generated from these files:
	config.proto

It has these top-level messages:
	Config
	NetworkConfig
	ChainConfig
	RPCConfig
	AppConfig
	PprofConfig
	MiscConfig
	StatsConfig
	InfluxdbConfig
	NbreConfig
*/
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
	return fileDescriptorConfig, []int{7, 0}
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
	App *AppConfig `protobuf:"bytes,102,opt,name=app" json:"app"`
	// Nbre Config.
	Nbre *NbreConfig `protobuf:"bytes,200,opt,name=nbre" json:"nbre"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

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

func (m *Config) GetNbre() *NbreConfig {
	if m != nil {
		return m.Nbre
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
	NetworkId            uint32 `protobuf:"varint,4,opt,name=network_id,json=networkId,proto3" json:"network_id"`
	StreamLimits         int32  `protobuf:"varint,5,opt,name=stream_limits,json=streamLimits,proto3" json:"stream_limits"`
	ReservedStreamLimits int32  `protobuf:"varint,6,opt,name=reserved_stream_limits,json=reservedStreamLimits,proto3" json:"reserved_stream_limits"`
}

func (m *NetworkConfig) Reset()                    { *m = NetworkConfig{} }
func (m *NetworkConfig) String() string            { return proto.CompactTextString(m) }
func (*NetworkConfig) ProtoMessage()               {}
func (*NetworkConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{1} }

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
}

func (m *ChainConfig) Reset()                    { *m = ChainConfig{} }
func (m *ChainConfig) String() string            { return proto.CompactTextString(m) }
func (*ChainConfig) ProtoMessage()               {}
func (*ChainConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{2} }

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
	HttpCors []string `protobuf:"bytes,6,rep,name=http_cors,json=httpCors" json:"http_cors"`
}

func (m *RPCConfig) Reset()                    { *m = RPCConfig{} }
func (m *RPCConfig) String() string            { return proto.CompactTextString(m) }
func (*RPCConfig) ProtoMessage()               {}
func (*RPCConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{3} }

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
	Pprof   *PprofConfig `protobuf:"bytes,6,opt,name=pprof" json:"pprof"`
	Version string       `protobuf:"bytes,100,opt,name=version,proto3" json:"version"`
}

func (m *AppConfig) Reset()                    { *m = AppConfig{} }
func (m *AppConfig) String() string            { return proto.CompactTextString(m) }
func (*AppConfig) ProtoMessage()               {}
func (*AppConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{4} }

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
	Memprofile string `protobuf:"bytes,3,opt,name=memprofile,proto3" json:"memprofile"`
}

func (m *PprofConfig) Reset()                    { *m = PprofConfig{} }
func (m *PprofConfig) String() string            { return proto.CompactTextString(m) }
func (*PprofConfig) ProtoMessage()               {}
func (*PprofConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{5} }

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
	DefaultKeystoreFileCiper string `protobuf:"bytes,1,opt,name=default_keystore_file_ciper,json=defaultKeystoreFileCiper,proto3" json:"default_keystore_file_ciper"`
}

func (m *MiscConfig) Reset()                    { *m = MiscConfig{} }
func (m *MiscConfig) String() string            { return proto.CompactTextString(m) }
func (*MiscConfig) ProtoMessage()               {}
func (*MiscConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{6} }

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
	Influxdb    *InfluxdbConfig `protobuf:"bytes,11,opt,name=influxdb" json:"influxdb"`
	MetricsTags []string        `protobuf:"bytes,12,rep,name=metrics_tags,json=metricsTags" json:"metrics_tags"`
}

func (m *StatsConfig) Reset()                    { *m = StatsConfig{} }
func (m *StatsConfig) String() string            { return proto.CompactTextString(m) }
func (*StatsConfig) ProtoMessage()               {}
func (*StatsConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{7} }

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
	Password string `protobuf:"bytes,5,opt,name=password,proto3" json:"password"`
}

func (m *InfluxdbConfig) Reset()                    { *m = InfluxdbConfig{} }
func (m *InfluxdbConfig) String() string            { return proto.CompactTextString(m) }
func (*InfluxdbConfig) ProtoMessage()               {}
func (*InfluxdbConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{8} }

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

type NbreConfig struct {
	// Nbre log path
	Log string `protobuf:"bytes,1,opt,name=log,proto3" json:"log"`
	// Nbre data path
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
	// Nbre runtime path
	Nbre string `protobuf:"bytes,3,opt,name=nbre,proto3" json:"nbre"`
	// Nbre admin address
	AdminAddress string `protobuf:"bytes,4,opt,name=admin_address,json=adminAddress,proto3" json:"admin_address"`
}

func (m *NbreConfig) Reset()                    { *m = NbreConfig{} }
func (m *NbreConfig) String() string            { return proto.CompactTextString(m) }
func (*NbreConfig) ProtoMessage()               {}
func (*NbreConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{9} }

func (m *NbreConfig) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

func (m *NbreConfig) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *NbreConfig) GetNbre() string {
	if m != nil {
		return m.Nbre
	}
	return ""
}

func (m *NbreConfig) GetAdminAddress() string {
	if m != nil {
		return m.AdminAddress
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
	proto.RegisterType((*NbreConfig)(nil), "nebletpb.NbreConfig")
	proto.RegisterEnum("nebletpb.StatsConfig_ReportingModule", StatsConfig_ReportingModule_name, StatsConfig_ReportingModule_value)
}

func init() { proto.RegisterFile("config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 1082 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x56, 0xdd, 0x6e, 0xdb, 0x46,
	0x13, 0xfd, 0x24, 0xcb, 0x8e, 0x38, 0x92, 0x1d, 0x65, 0xe3, 0x38, 0x9b, 0xf8, 0xab, 0xad, 0xaa,
	0x30, 0xa0, 0x22, 0x85, 0x8b, 0xba, 0xb9, 0xe9, 0x45, 0x2f, 0x0c, 0x01, 0x05, 0x0c, 0xff, 0xc0,
	0xa0, 0xdb, 0x6b, 0x82, 0x22, 0x47, 0xd4, 0xc2, 0x14, 0x97, 0xd8, 0x5d, 0x39, 0xf1, 0x5d, 0xdf,
	0xa4, 0xaf, 0xd1, 0x47, 0xe8, 0x13, 0xb4, 0x4f, 0x53, 0xa0, 0x98, 0xe1, 0x52, 0x94, 0x84, 0xdc,
	0xed, 0x9c, 0x73, 0x76, 0x87, 0x9a, 0x39, 0x3b, 0x2b, 0xe8, 0x27, 0xba, 0x98, 0xa9, 0xec, 0xbc,
	0x34, 0xda, 0x69, 0xd1, 0x2d, 0x70, 0x9a, 0xa3, 0x2b, 0xa7, 0xa3, 0x3f, 0xdb, 0xb0, 0x37, 0x61,
	0x4a, 0xfc, 0x00, 0x2f, 0x0a, 0x74, 0x9f, 0xb4, 0x79, 0x94, 0xad, 0x61, 0x6b, 0xdc, 0xbb, 0x78,
	0x7b, 0x5e, 0xcb, 0xce, 0xef, 0x2a, 0xa2, 0x52, 0x86, 0xb5, 0x4e, 0x7c, 0x80, 0xdd, 0x64, 0x1e,
	0xab, 0x42, 0xb6, 0x79, 0xc3, 0x9b, 0x66, 0xc3, 0x84, 0x60, 0x2f, 0xaf, 0x34, 0xe2, 0x0c, 0x76,
	0x4c, 0x99, 0xc8, 0x1d, 0x96, 0xbe, 0x6e, 0xa4, 0xe1, 0xfd, 0xc4, 0x0b, 0x89, 0xa7, 0x33, 0xad,
	0x8b, 0x9d, 0x95, 0xe9, 0xf6, 0x99, 0x0f, 0x04, 0xd7, 0x67, 0xb2, 0x46, 0x8c, 0xa1, 0xb3, 0x50,
	0x36, 0x91, 0xc8, 0xda, 0xc3, 0x46, 0x7b, 0xab, 0x6c, 0xe2, 0xa5, 0xac, 0xa0, 0xec, 0x71, 0x59,
	0xca, 0xd9, 0x76, 0xf6, 0xcb, 0xb2, 0xac, 0xb3, 0xc7, 0x65, 0x29, 0xbe, 0x85, 0x4e, 0x31, 0x35,
	0x28, 0xff, 0x6a, 0x6d, 0x9f, 0x78, 0x37, 0x35, 0x58, 0x9f, 0x48, 0x92, 0xd1, 0xdf, 0x2d, 0xd8,
	0xdf, 0xa8, 0x8b, 0x10, 0xd0, 0xb1, 0x88, 0xa9, 0x6c, 0x0d, 0x77, 0xc6, 0x41, 0xc8, 0x6b, 0x71,
	0x04, 0x7b, 0xb9, 0xb2, 0x0e, 0xa9, 0x46, 0x84, 0xfa, 0x48, 0x9c, 0x42, 0xaf, 0x34, 0xea, 0x29,
	0x76, 0x18, 0x3d, 0xe2, 0x33, 0x57, 0x25, 0x08, 0xc1, 0x43, 0xd7, 0xf8, 0x2c, 0xbe, 0x02, 0xf0,
	0x65, 0x8e, 0x54, 0x2a, 0x3b, 0xc3, 0xd6, 0x78, 0x3f, 0x0c, 0x3c, 0x72, 0x95, 0x8a, 0x6f, 0x60,
	0xdf, 0x3a, 0x83, 0xf1, 0x22, 0xca, 0xd5, 0x42, 0x39, 0x2b, 0x77, 0x87, 0xad, 0xf1, 0x6e, 0xd8,
	0xaf, 0xc0, 0x1b, 0xc6, 0xc4, 0x47, 0x38, 0x32, 0x68, 0xd1, 0x3c, 0x61, 0x1a, 0x6d, 0xaa, 0xf7,
	0x58, 0x7d, 0x58, 0xb3, 0x0f, 0x6b, 0xbb, 0x46, 0x7f, 0x74, 0xa0, 0xb7, 0xd6, 0x3f, 0xf1, 0x0e,
	0xba, 0xdc, 0x41, 0xfa, 0x8e, 0x16, 0x7f, 0xc7, 0x0b, 0x8e, 0xaf, 0x52, 0x21, 0xe1, 0x45, 0x86,
	0x05, 0x5a, 0x65, 0xd9, 0x02, 0x41, 0x58, 0x87, 0xc4, 0xa4, 0xb1, 0x8b, 0x53, 0x65, 0x64, 0xaf,
	0x62, 0x7c, 0x48, 0x15, 0x79, 0xc4, 0x67, 0x22, 0xfa, 0x4c, 0xf8, 0x88, 0x7e, 0xb0, 0x75, 0xb1,
	0x71, 0xd1, 0x42, 0x15, 0x28, 0x0f, 0x87, 0xad, 0x71, 0x37, 0x0c, 0x18, 0xb9, 0x55, 0x05, 0x8a,
	0xf7, 0xd0, 0x4d, 0xb4, 0x2a, 0xa6, 0xb1, 0x45, 0xf9, 0x86, 0x37, 0xae, 0x62, 0x71, 0x08, 0xbb,
	0xb4, 0xc9, 0xc8, 0x23, 0x26, 0xaa, 0x40, 0x9c, 0x00, 0x94, 0xb1, 0xb5, 0xe5, 0xdc, 0xd0, 0x9e,
	0xb7, 0xbe, 0xc2, 0x2b, 0x44, 0xfc, 0x04, 0xef, 0xb0, 0x88, 0xa7, 0x39, 0x46, 0x06, 0x17, 0xda,
	0x61, 0x64, 0x55, 0x56, 0x44, 0x5c, 0x10, 0x23, 0x25, 0xe7, 0x3f, 0xaa, 0x04, 0x21, 0xf3, 0x0f,
	0x2a, 0x2b, 0x1e, 0x98, 0x15, 0xdf, 0x81, 0xf8, 0xc2, 0x9e, 0x77, 0x9c, 0x62, 0x60, 0xb6, 0xd5,
	0xc7, 0x10, 0x64, 0xb1, 0x8d, 0x4a, 0xa3, 0x12, 0x94, 0xef, 0xab, 0x6f, 0xcf, 0x62, 0x7b, 0x4f,
	0x71, 0x4d, 0x72, 0x5f, 0xe4, 0xf1, 0x8a, 0xe4, 0x5e, 0x88, 0x0f, 0xf0, 0x8a, 0x12, 0xc4, 0x6e,
	0x69, 0x30, 0x4a, 0x54, 0x39, 0x47, 0x63, 0xe5, 0xff, 0xd9, 0x48, 0x83, 0x15, 0x31, 0xa9, 0x70,
	0x2e, 0xe0, 0xb2, 0x44, 0x13, 0x15, 0x3a, 0x45, 0x79, 0xe2, 0x0b, 0x48, 0xc8, 0x9d, 0x4e, 0x51,
	0x7c, 0x0f, 0xaf, 0x97, 0x85, 0x5d, 0x96, 0xa5, 0x36, 0x0e, 0x53, 0x72, 0xdd, 0x27, 0x6d, 0x52,
	0x79, 0xca, 0x29, 0xc5, 0x1a, 0x75, 0x5d, 0x31, 0xdc, 0xc2, 0xe7, 0x22, 0xb6, 0xee, 0x59, 0x0e,
	0x7d, 0x0b, 0xab, 0x70, 0xf4, 0x4f, 0x0b, 0x82, 0xd5, 0xb5, 0xa5, 0xbc, 0xa6, 0x4c, 0x22, 0x6f,
	0xf3, 0xca, 0xfc, 0x81, 0x29, 0x93, 0x9b, 0x95, 0xd3, 0xe7, 0xce, 0x95, 0xd1, 0xc6, 0x35, 0x00,
	0x82, 0xb6, 0x04, 0x0b, 0x9d, 0x2e, 0x73, 0x94, 0x3b, 0x8d, 0xe0, 0x96, 0x11, 0xaa, 0x42, 0xa2,
	0x8b, 0x02, 0x13, 0xa7, 0x74, 0x51, 0x3b, 0xb8, 0xc3, 0x0e, 0x1e, 0x34, 0x84, 0xf7, 0x7c, 0x93,
	0x6e, 0xed, 0x5a, 0xf8, 0x74, 0x2c, 0x38, 0x86, 0x80, 0x05, 0x89, 0x36, 0x74, 0x0f, 0x28, 0x59,
	0x97, 0x80, 0x89, 0x36, 0x76, 0xf4, 0x6f, 0x0b, 0x82, 0xd5, 0x48, 0x20, 0x69, 0xae, 0xb3, 0x28,
	0xc7, 0x27, 0xcc, 0xd9, 0xfa, 0x41, 0xd8, 0xcd, 0x75, 0x76, 0x43, 0x31, 0x5d, 0x0b, 0x22, 0x67,
	0x2a, 0xc7, 0xda, 0xfc, 0xb9, 0xce, 0x7e, 0x51, 0x39, 0x8a, 0xb7, 0x40, 0xcb, 0x28, 0xce, 0x90,
	0x2f, 0xf6, 0x7e, 0xb8, 0x97, 0xeb, 0xec, 0x32, 0x43, 0x71, 0x0e, 0xaf, 0xbd, 0xe5, 0x12, 0x13,
	0xdb, 0x79, 0x64, 0x90, 0x4a, 0xce, 0xbf, 0xa5, 0x1b, 0xbe, 0xaa, 0xa8, 0x09, 0x31, 0x21, 0x13,
	0x62, 0x0c, 0x83, 0x75, 0x61, 0xb4, 0x34, 0x39, 0xff, 0xa2, 0x20, 0x3c, 0x48, 0x1a, 0xd9, 0x6f,
	0x26, 0xa7, 0xb1, 0x59, 0x96, 0x46, 0xcf, 0xf8, 0x66, 0x6f, 0x8c, 0xcd, 0x7b, 0x82, 0xeb, 0xb1,
	0xc9, 0x1a, 0xea, 0xec, 0x13, 0x1a, 0xab, 0x74, 0xc1, 0x53, 0x36, 0x08, 0xeb, 0x70, 0x54, 0x40,
	0x6f, 0x4d, 0xbf, 0xdd, 0xbb, 0xaa, 0x04, 0xeb, 0xbd, 0x3b, 0x01, 0x48, 0xca, 0x25, 0xed, 0x68,
	0xca, 0xb0, 0x86, 0x10, 0xbf, 0xc0, 0x45, 0xcd, 0xfb, 0x29, 0xd7, 0x20, 0xa3, 0x6b, 0x80, 0x66,
	0x54, 0x8b, 0x9f, 0xe1, 0x38, 0xc5, 0x59, 0xbc, 0xcc, 0x1d, 0xd9, 0xd3, 0x3a, 0x6d, 0x90, 0xeb,
	0x4b, 0xd6, 0x47, 0xe3, 0xd3, 0x4b, 0x2f, 0xb9, 0xf6, 0x0a, 0xaa, 0xf8, 0x84, 0xf8, 0xd1, 0xef,
	0x6d, 0xe8, 0xad, 0x3d, 0x12, 0xe2, 0x0c, 0x0e, 0x7c, 0xb5, 0x17, 0xe8, 0x8c, 0x4a, 0x2c, 0x9f,
	0xd0, 0x0d, 0xf7, 0x2b, 0xf4, 0xb6, 0x02, 0xc5, 0x3d, 0x0c, 0xaa, 0xf2, 0xaa, 0x22, 0xab, 0x4d,
	0x48, 0x2e, 0x3d, 0xb8, 0x38, 0xfb, 0xe2, 0xe3, 0x73, 0x1e, 0xd6, 0xea, 0xca, 0x9f, 0xe1, 0x4b,
	0xb3, 0x09, 0x88, 0x8f, 0xd0, 0x55, 0xc5, 0x2c, 0x5f, 0x7e, 0x4e, 0xa7, 0x3c, 0xfd, 0x7a, 0x17,
	0xb2, 0x39, 0xe9, 0xca, 0x33, 0xbe, 0x25, 0x2b, 0xa5, 0xf8, 0x1a, 0xfa, 0xfe, 0x3b, 0x23, 0x17,
	0x67, 0x56, 0xf6, 0xd9, 0x9b, 0x3d, 0x8f, 0xfd, 0x1a, 0x67, 0x76, 0x74, 0x0a, 0x2f, 0xb7, 0x92,
	0x8b, 0x3e, 0x74, 0xeb, 0x13, 0x07, 0xff, 0x1b, 0x7d, 0x86, 0x83, 0xcd, 0xf3, 0xe9, 0x51, 0x9a,
	0x6b, 0xeb, 0x7c, 0xf1, 0x78, 0x4d, 0x18, 0xfb, 0xae, 0xcd, 0xe6, 0xe4, 0xb5, 0x38, 0x80, 0x76,
	0x3a, 0xf5, 0x1d, 0x6a, 0xa7, 0x53, 0xd2, 0x2c, 0x2d, 0x1a, 0xf6, 0x66, 0x10, 0xf2, 0x9a, 0x66,
	0x30, 0xcd, 0x4f, 0x9e, 0x1b, 0x95, 0x0d, 0x57, 0xf1, 0xe8, 0x11, 0xa0, 0x79, 0x22, 0xc5, 0x00,
	0x76, 0x72, 0x9d, 0xf9, 0xa4, 0xb4, 0xa4, 0xf3, 0xe8, 0x05, 0xf0, 0x1e, 0xe1, 0x35, 0x61, 0xfc,
	0xda, 0x56, 0x59, 0x79, 0x4d, 0x0f, 0x5b, 0x9c, 0x2e, 0x54, 0x11, 0xc5, 0x69, 0x6a, 0xd0, 0x5a,
	0xff, 0x01, 0x7d, 0x06, 0x2f, 0x2b, 0x6c, 0xba, 0xc7, 0xff, 0x63, 0x7e, 0xfc, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0x47, 0x6f, 0x5e, 0xac, 0xd7, 0x08, 0x00, 0x00,
}
