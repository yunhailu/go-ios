package diagnostics

import ios "github.com/danielpaulus/go-ios/ios"

func ioregentryRequest(key string) []byte {
	return ioregentryRequestBase(key, "EntryName")
	// requestMap := map[string]interface{}{
	// 	"Request":   "IORegistry",
	// 	"EntryName": key,
	// }
	// bt, err := ios.PlistCodec{}.Encode(requestMap)
	// if err != nil {
	// 	panic("query request encoding should never fail")
	// }
}

func ioregentryRequestBase(key string, mType string) []byte {
	requestMap := map[string]interface{}{
		"Request": "IORegistry",
		mType:     key,
	}
	bt, err := ios.PlistCodec{}.Encode(requestMap)
	if err != nil {
		panic("gestalt request encoding should never fail")
	}
	return bt
}

func (diagnosticsConn *Connection) IORegistry(key string, mType string) ([]byte, error) {
	err := diagnosticsConn.deviceConn.Send(ioregentryRequestBase(key, mType))
	if err != nil {
		return nil, err
	}
	respBytes, err := diagnosticsConn.plistCodec.Decode(diagnosticsConn.deviceConn.Reader())
	if err != nil {
		return nil, err
	}
	// err = diagnosticsConn.deviceConn.Send(goodBye())
	// if err != nil {
	// 	return nil, err
	// }
	return respBytes, nil

}

func (diagnosticsConn *Connection) IORegEntryQuery(key string) (interface{}, error) {
	err := diagnosticsConn.deviceConn.Send(ioregentryRequest(key))
	if err != nil {
		return "", err
	}
	respBytes, err := diagnosticsConn.plistCodec.Decode(diagnosticsConn.deviceConn.Reader())
	if err != nil {
		return "", err
	}
	plist, err := ios.ParsePlist(respBytes)
	return plist, err
}
