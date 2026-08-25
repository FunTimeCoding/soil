package macos

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/join"
	"strings"
)

func CreateBundle(
	name string,
	path string,
	executablePath string,
	iconPath string,
	vendor string,
	version string,
) {
	lower := strings.ToLower(name)
	identifier := fmt.Sprintf("%s.%s", vendor, lower)
	content := join.Absolute(
		join.Absolute(
			system.WorkDirectory(),
			path,
			fmt.Sprintf("%s.app", name),
		),
		"Contents",
	)
	macos := join.Absolute(content, "MacOS")
	resource := join.Absolute(content, "Resources")
	executable := join.Absolute(macos, lower)
	system.MakeDirectory(macos)
	system.MakeDirectory(resource)
	system.CopyFile(executablePath, executable)
	system.Executable(executable)
	system.CopyFile(iconPath, join.Absolute(resource, "icon.icns"))
	system.SaveFile(
		join.Absolute(content, "Info.plist"),
		fmt.Sprintf(
			`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" 
    "https://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>CFBundleName</key>
    <string>%s</string>

    <key>CFBundleDisplayName</key>
    <string>%sDisplayName</string>

    <key>CFBundleIdentifier</key>
    <string>%s</string>

    <key>CFBundleVersion</key>
    <string>%s</string>

    <key>CFBundleShortVersionString</key>
    <string>%s</string>

    <key>CFBundleExecutable</key>
    <string>%s</string>

    <key>CFBundlePackageType</key>
    <string>APPL</string>

    <key>LSUIElement</key>
    <true/>

    <key>CFBundleIconFile</key>
    <string>icon</string>
</dict>
</plist>
`,
			name,
			name,
			identifier,
			version,
			version,
			lower,
		),
	)
}
