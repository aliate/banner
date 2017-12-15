package banner_test

import (
	. "github.com/aliate/banner"
	"testing"
)

func showBanners(adjoin bool) {
	NewBanner("Monday", adjoin).Show()
	NewBanner("Tuseday", adjoin).Show()
	NewBanner("Thursday", adjoin).Show()
	NewBanner("Friday", adjoin).Show()
	NewBanner("Saturday", adjoin).Show()
	NewBanner("Sunday", adjoin).Show()
	NewBanner("Winter", adjoin).Show()
	NewBanner("Autum", adjoin).Show()
	NewBanner("jump", adjoin).Show()
	NewBanner("Jump", adjoin).Show()
	NewBanner("golang", adjoin).Show()
	NewBanner("Golang", adjoin).Show()
	NewBanner("Echo", adjoin).Show()
	NewBanner("Autopattern", adjoin).Show()
	NewBanner("Spring", adjoin).Show()
}

func TestNewBanner_echo(t *testing.T) {
    showBanners(false)
    showBanners(true)
}

