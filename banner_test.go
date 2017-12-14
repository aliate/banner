package banner_test

import (
	"testing"
	. "github.com/aliate/banner"
)


func TestNewBanner_1(t *testing.T) { NewBanner("Monday").Show() }
func TestNewBanner_2(t *testing.T) { NewBanner("Tuseday").Show() }
func TestNewBanner_3(t *testing.T) { NewBanner("Wednesday").Show() }
func TestNewBanner_4(t *testing.T) { NewBanner("Thursday").Show() }
func TestNewBanner_5(t *testing.T) { NewBanner("Friday").Show() }
func TestNewBanner_6(t *testing.T) { NewBanner("Saturday").Show() }
func TestNewBanner_7(t *testing.T) { NewBanner("Sunday").Show() }


func TestNewBanner_8(t *testing.T) { NewBanner("zoom").Show() }
func TestNewBanner_9(t *testing.T) { NewBanner("Zirjom").Show() }


func TestNewBanner_10(t *testing.T) { NewBanner("Water").Show() }
func TestNewBanner_11(t *testing.T) { NewBanner("Spring").Show() }
func TestNewBanner_12(t *testing.T) { NewBanner("Summer").Show() }
func TestNewBanner_13(t *testing.T) { NewBanner("Autumn").Show() }

func TestNewBanner_echo(t *testing.T) {
	NewBanner("Echo").Show()
	NewBanner("Autopattern").Show()
}

//func TestNewBanner_5(t *testing.T) { NewBanner("Friday").Show() }
//func TestNewBanner_11(t *testing.T) { NewBanner("Spring").Show() }
