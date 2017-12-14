package banner_test

import (
	"testing"
	. "github.com/aliate/banner"
)

func TestGetBanner_1(t *testing.T) { GetBanner("Monday").Show() }
func TestGetBanner_2(t *testing.T) { GetBanner("Tuseday").Show() }
func TestGetBanner_3(t *testing.T) { GetBanner("Wednesday").Show() }
func TestGetBanner_4(t *testing.T) { GetBanner("Thursday").Show() }
func TestGetBanner_5(t *testing.T) { GetBanner("Friday").Show() }
func TestGetBanner_6(t *testing.T) { GetBanner("Saturday").Show() }
func TestGetBanner_7(t *testing.T) { GetBanner("Sunday").Show() }


func TestGetBanner_8(t *testing.T) { GetBanner("zoom").Show() }
func TestGetBanner_9(t *testing.T) { GetBanner("Ziroom").Show() }


func TestGetBanner_10(t *testing.T) { GetBanner("Water").Show() }
func TestGetBanner_11(t *testing.T) { GetBanner("Spring").Show() }
func TestGetBanner_12(t *testing.T) { GetBanner("Summer").Show() }
func TestGetBanner_13(t *testing.T) { GetBanner("Autumn").Show() }

