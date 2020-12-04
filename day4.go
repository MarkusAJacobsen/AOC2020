package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	Byr int
	Iyr int
	Eyr int
	Hgt string
	Hcl string
	Ecl string
	Pid int
	Cid string
}

func Day4(in []string) int {
	var candPassports [][]string

	start := 0
	offset := 0
	for i, ps := range in {
		if ps != "" {
			offset++
		} else if ps == "" {
			candPassports = append(candPassports, in[start:start+offset])
			start = i + 1
			offset = 0
		}
	}

	var passports []Passport
	for _, ps := range candPassports {
		psFlat := ""
		for _, psParts := range ps {
			psFlat = psFlat + " " + psParts
		}

		parts := strings.Split(psFlat, " ")

		var passport Passport
		for _, p := range parts {
			if p == "" {
				continue
			}
			switch p[0:3] {
			case "byr":
				byr, _ := strconv.Atoi(p[4:])

				if byr >= 1920 && byr <= 2002 {
					passport.Byr = byr
				}
			case "iyr":
				iyr, _ := strconv.Atoi(p[4:])

				if iyr >= 2010 && iyr <= 2020 {
					passport.Iyr = iyr
				}
			case "eyr":
				eyr, _ := strconv.Atoi(p[4:])

				if eyr >= 2020 && eyr <= 2030 {
					passport.Eyr = eyr
				}
			case "hgt":
				re := regexp.MustCompile("[a-z]+")
				split := re.Split(p[4:], -1)
				hgt, _ := strconv.Atoi(split[0])
				isCm, _ := regexp.MatchString(`(?m)cm`, p[4:])
				isIn, _ := regexp.MatchString(`(?m)in`, p[4:])

				if hgt >= 150 && hgt <= 193 && isCm {
					passport.Hgt = p[4:]
				} else if hgt >= 59 && hgt <= 76 && isIn {
					passport.Hgt = p[4:]
				}
			case "hcl":
				rgb, _ := regexp.MatchString("^#([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$", p[4:])

				if rgb {
					passport.Hcl = p[4:]
				}
			case "ecl":
				ecl := p[4:]

				if ecl == "amb" || ecl == "blu" || ecl == "brn" || ecl == "gry" || ecl == "grn" || ecl == "hzl" || ecl == "oth" {
					passport.Ecl = ecl
				}
			case "pid":
				pid, _ := regexp.MatchString("[0-9]{9}", p[4:])

				if pid {
					if len(p[4:]) == 9 {
						passport.Pid, _ = strconv.Atoi(p[4:])
					}
				}
			case "cid":
				passport.Cid = p[4:]
			}
		}
		passports = append(passports, passport)
	}

	validPassports := getValidPassportsAndNorthPoleId(passports)
	fmt.Println(len(validPassports))
	return len(validPassports)
}

func getValidPassportsAndNorthPoleId(ps []Passport) []Passport {
	var valid []Passport

	for _, p := range ps {
		if p.Byr != 0 && p.Iyr != 0 && p.Eyr != 0 && p.Hgt != "" && p.Hcl != "" && p.Ecl != "" && p.Pid != 0 {
			valid = append(valid, p)
		}
	}

	return valid
}
