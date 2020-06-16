// У меня все зафейлено: Main должен быть из пары строк, а остальное - разнесено
// по функциям, чтобы можно было тестировать эти функции
package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
