package aoc

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/jaynak/aoc2020/pkg/util"
)

func Day4(path string) (int, int) {

	batch := util.ReadToStrings(path)

	var passports []map[string]string

	passport := make(map[string]string)
	for _, line := range batch {
		if line == "" {
			passports = append(passports, passport)
			passport = make(map[string]string)
		} else {
			// Split the line into key/value pairs
			kv := strings.Split(line, " ")
			for _, kvpair := range kv {
				entry := strings.Split(kvpair, ":")
				passport[entry[0]] = entry[1]
			}
		}
	}

	// Add the last passport
	passports = append(passports, passport)
	validPassports := 0
	reallyValid := 0
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, passport := range passports {
		valid := true
		for _, field := range requiredFields {
			if _, ok := passport[field]; !ok {
				valid = false
			}
		}

		if valid {

			// Check all the other rules here
			if IsValidPassport(passport) {
				reallyValid++
			}
			validPassports++
		}
	}

	return validPassports, reallyValid
}

func IsValidPassport(passport map[string]string) bool {

	heightRegex := regexp.MustCompile("^([0-9]+)([cmin]{2})$")
	hairRegex := regexp.MustCompile("^#[0-9a-f]{6}$")
	passportNumberRegex := regexp.MustCompile("^[0-9]{9}$")

	eyeColour := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	for k, v := range passport {
		switch k {
		case "byr":
			yr, err := strconv.Atoi(v)
			if err != nil || yr < 1920 || yr > 2002 {
				return false
			}
		case "iyr":
			yr, err := strconv.Atoi(v)
			if err != nil || yr < 2010 || yr > 2020 {
				return false
			}
		case "eyr":
			yr, err := strconv.Atoi(v)
			if err != nil || yr < 2020 || yr > 2030 {
				return false
			}
		case "hgt":
			match := heightRegex.FindAllStringSubmatch(v, -1)
			if len(match) == 0 {
				return false
			} else {
				h, err := strconv.Atoi(match[0][1])
				if err != nil {
					return false
				}

				switch match[0][2] {
				case "cm":
					if h < 150 || h > 193 {
						return false
					}
				case "in":
					if h < 59 || h > 76 {
						return false
					}
				default:
					return false
				}
			}
		case "hcl":
			if !hairRegex.MatchString(v) {
				return false
			}
		case "ecl":
			found := false
			for _, eyeCol := range eyeColour {
				if v == eyeCol {
					found = true
				}
			}

			if !found {
				return false
			}
		case "pid":
			if !passportNumberRegex.MatchString(v) {
				return false
			}
		}
	}

	// Fell through
	return true
}
