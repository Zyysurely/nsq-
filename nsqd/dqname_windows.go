// +build windows

package nsqd

// windows缓减下的磁盘存储名
// On Windows, file names cannot contain colons.
func getBackendName(topicName, channelName string) string {
	// backend names, for uniqueness, automatically include the topic... <topic>;<channel>
	backendName := topicName + ";" + channelName
	return backendName
}
