package storage

import (
	"github.com/ectrc/snow/aid"
)

func GetDefaultEngine() []byte {
	return []byte(`
[OnlineSubsystemMcp.Xmpp]
bUseSSL=false
Protocol=ws
ServerAddr="ws://`+ aid.Config.API.Host + aid.Config.API.Port +`"

[OnlineSubsystemMcp.Xmpp Prod]
bUseSSL=false
Protocol=ws
ServerAddr="ws://`+ aid.Config.API.Host + aid.Config.API.Port +`"

[OnlineSubsystemMcp]
bUsePartySystemV2=false

[OnlineSubsystemMcp.OnlinePartySystemMcpAdapter]
bUsePartySystemV2=false

[XMPP]
bEnableWebsockets=true

[ConsoleVariables]
n.VerifyPeer=0
FortMatchmakingV2.ContentBeaconFailureCancelsMatchmaking=0
Fort.ShutdownWhenContentBeaconFails=0
FortMatchmakingV2.EnableContentBeacon=0

[/Script/Qos.QosRegionManager]
NumTestsPerRegion=5
PingTimeout=3.0`)
}