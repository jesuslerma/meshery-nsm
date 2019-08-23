package nsm

type supportedOperation struct {
	// a friendly name
	name string
	// the template file name
	templateName string
	// the app label
	appLabel string
	// // returnLogs specifies if the operation logs should be returned
	// returnLogs bool
}

const (
	customOpCommand    = "custom"
	installNSMCommand  = "nsm_install"
	installICMPCommand = "icmp_install"
)

var supportedOps = map[string]supportedOperation{
	installNSMCommand: {
		name: "Install Network Service Mesh",
	},
	installICMPCommand: {
		name: "Install ICMP Application",
	},
	customOpCommand: {
		name: "Custom YAML",
	},
}
