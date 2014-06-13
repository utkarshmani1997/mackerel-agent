package metrics

import (
	"strconv"
	"strings"

	"github.com/mackerelio/mackerel-agent/config"
	"github.com/mackerelio/mackerel-agent/logging"
	"github.com/mackerelio/mackerel-agent/util"
)

type PluginGenerator struct {
	Config config.PluginConfig
}

var pluginLogger = logging.GetLogger("metrics.plugin")

const pluginPrefix = "custom."

func (g *PluginGenerator) Generate() (Values, error) {
	command := g.Config.Command

	pluginLogger.Debugf(`Executing plugin: command = "%s"`, command)

	stdout, stderr, err := util.RunCommand(command)

	if err != nil {
		pluginLogger.Errorf(`Failed to execute command "%s" (skip these metrics):\n%s`, command, stderr)
		return nil, err
	}

	results := make(map[string]float64, 0)
	for _, line := range strings.Split(stdout, "\n") {
		if line == "" {
			continue
		}

		// Key, value, timestamp
		// ex.) localhost.localdomain.tcp.CLOSING 0 1397031808
		items := strings.Split(line, "\t")
		if len(items) != 3 {
			pluginLogger.Warningf("Output line malformed: %s", line)
			continue
		}

		value, err := strconv.ParseFloat(items[1], 64)
		if err != nil {
			pluginLogger.Warningf("Failed to parse metric value: %s", err)
			continue
		}

		results[pluginPrefix+items[0]] = value
	}

	return results, nil
}
