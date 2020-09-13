package netconf

import (
	"fmt"
	"strconv"
)

// H3CMethodAction define action operation to h3c device
func H3CMethodAction(action string) RawMethod{
	return RawMethod(fmt.Sprintf("<action>%s</action>", action))
}

// H3CMethodCLI define cli operation to h3c device
func H3CMethodCLI(cli string) RawMethod{
	return RawMethod(fmt.Sprintf("<CLI>%s</CLI>", cli))
}

// H3CMethodEditConfig define edit config to h3c device
func H3CMethodEditConfig(target string, config string) RawMethod{
	return RawMethod(fmt.Sprintf("<edit-config><target><%s/></target>%s</edit-config>", target, config))
}
// H3CMethodGet define get h3c device data
func H3CMethodGet(filter string) RawMethod{
	return RawMethod(fmt.Sprintf("<get>%s</get>", filter))
}

// H3CMethodGetBulk define get h3c device data
func H3CMethodGetBulk(filter string) RawMethod{
	return RawMethod(fmt.Sprintf("<get-bulk>%s</get-bulk>", filter))
}

// H3CMethodGetBulkConfig define get bulk config from h3c device
func H3CMethodGetBulkConfig(target string, config string) RawMethod{
	return RawMethod(fmt.Sprintf("<get-bulk-config><target><%s/></target>%s</get-bulk-config>", target, config))
}

// H3CMethodGetConfig define get config from h3c device
func H3CMethodGetConfig(target string, config string) RawMethod{
	return RawMethod(fmt.Sprintf("<get-config><target><%s/></target>%s</get-config>", target, config))
}

// H3CMethodLoad define load a file merge to h3c device
func H3CMethodLoad(name string) RawMethod{
	return RawMethod(fmt.Sprintf("<load><file>%s</file></load>", name)) 
}

// H3CMethodLock define lock the operation to h3c device
func H3CMethodLock(target string) RawMethod{
	return RawMethod(fmt.Sprintf("<lock><target><%s/></target></lock>", target)) 
}

// H3CMethodRollback define use a file.cfg to rollback the h3c device
func H3CMethodRollback(name string) RawMethod{
	return RawMethod(fmt.Sprintf("<rollback><file>%s</file></rollback>", name)) 
}

// H3CMethodSave define use a file.cfg to save the h3c device
func H3CMethodSave(overwrite bool, binary bool, name string) RawMethod{
	return RawMethod(fmt.Sprintf("<save OverWrite=\"%s\" Binary-only=\"%s\"><file>%s</file></save>", strconv.FormatBool(overwrite), strconv.FormatBool(binary), name)) 
}
// H3CMethodUnLock define unlock the operation to h3c device
func H3CMethodUnLock(target string) RawMethod{
	return RawMethod(fmt.Sprintf("<unlock><target><%s/></target></unlock>", target)) 
}