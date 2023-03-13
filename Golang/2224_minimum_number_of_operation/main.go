package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertTime(current string, correct string) int {
  // validate if current or correct is not empty
  if current == "" || correct == "" {
    return 0
  }

  // curr_split[0] for the current hour
  // curr_split[1] for the current minutes
  curr_split := strings.Split(current, ":")
  curr_hour, err_ch := strconv.Atoi(curr_split[0])
  curr_minutes, err_cm := strconv.Atoi(curr_split[1])
  if err_ch != nil || err_cm != nil {
    return 0
  }

  total_curr_min := curr_minutes + (curr_hour * 60)

  // corr_split[0] for the correct hour
  // corr_split[1] for the correct minutes
  corr_split := strings.Split(correct, ":")
  corr_hour, err_ch := strconv.Atoi(corr_split[0])
  corr_minutes, err_cm := strconv.Atoi(corr_split[1])
  if err_ch != nil || err_cm != nil {
    return 0
  }

  total_corr_min := corr_minutes + (corr_hour * 60)

  var result int
  minutes_left := total_corr_min - total_curr_min

  for minutes_left != 0 {
    if minutes_left >= 60 {
      minutes_left -= 60
    } else if minutes_left >= 15 {
      minutes_left -= 15
    } else if minutes_left >= 5 {
      minutes_left -= 5
    } else {
      minutes_left -= 1
    }

    result++
  }


  return result
}

func main() {
  fmt.Println(convertTime("11:00", "11:01"))
}
