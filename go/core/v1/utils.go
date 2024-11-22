package "github.com/sourav-py/klusterduck/go/core/v1/utils"


func matchLabels(selector map[string]string, labels map[string]string) bool {
    for key, value := range selector {
        if podVal, ok := labels[key]; !ok || podVal != value {
            return false
        }
    }
    return true
}

