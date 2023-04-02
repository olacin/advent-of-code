use anyhow::Result;

#[derive(Debug)]
enum HeightUnit {
    None,
    Centimeters,
    Inches,
}

#[derive(Debug)]
struct Height {
    value: usize,
    unit: HeightUnit,
}

impl TryFrom<String> for Height {
    type Error = anyhow::Error;

    fn try_from(value: String) -> Result<Self, Self::Error> {
        if value.ends_with("cm") {
            let h = value.replace("cm", "").parse::<usize>()?;
            return Ok(Height {
                value: h,
                unit: HeightUnit::Centimeters,
            });
        } else if value.ends_with("in") {
            let h = value.replace("in", "").parse::<usize>()?;
            return Ok(Height {
                value: h,
                unit: HeightUnit::Inches,
            });
        }
        Ok(Height {
            value: value.parse::<usize>()?,
            unit: HeightUnit::None,
        })
    }
}

impl Height {
    fn is_valid(&self) -> bool {
        match self.unit {
            HeightUnit::Centimeters => self.value >= 150 && self.value <= 193,
            HeightUnit::Inches => self.value >= 59 && self.value <= 76,
            HeightUnit::None => false,
        }
    }
}

#[derive(Default, Debug)]
struct Passport {
    birth_year: Option<usize>,
    issue_year: Option<usize>,
    expiration_year: Option<usize>,
    height: Option<Height>,
    hair_color: Option<String>,
    eye_color: Option<String>,
    passport_id: Option<String>,
    country_id: Option<u16>,
}

impl TryFrom<String> for Passport {
    type Error = anyhow::Error;

    fn try_from(value: String) -> Result<Self, Self::Error> {
        let mut passport = Passport::default();
        for pt in value.split(" ") {
            if let Some((key, val)) = pt.split_once(":") {
                match key {
                    "byr" => passport.birth_year = val.parse::<usize>().ok(),
                    "iyr" => passport.issue_year = val.parse::<usize>().ok(),
                    "eyr" => passport.expiration_year = val.parse::<usize>().ok(),
                    "hgt" => passport.height = Height::try_from(val.to_string()).ok(),
                    "hcl" => passport.hair_color = Some(val.to_string()),
                    "ecl" => passport.eye_color = Some(val.to_string()),
                    "pid" => passport.passport_id = Some(val.to_string()),
                    "cid" => passport.country_id = val.parse::<u16>().ok(),
                    _ => {}
                }
            }
        }
        return Ok(passport);
    }
}

impl Passport {
    fn has_values(&self) -> bool {
        self.birth_year.is_some()
            && self.issue_year.is_some()
            && self.expiration_year.is_some()
            && self.height.is_some()
            && self.hair_color.is_some()
            && self.eye_color.is_some()
            && self.passport_id.is_some()
    }

    fn is_valid(&self) -> bool {
        let eye_colors: [&str; 7] = ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"];
        // Birth year
        self.birth_year.is_some()
            && self.birth_year.unwrap() >= 1920
            && self.birth_year.unwrap() <= 2002
            // Issue year
            && self.issue_year.is_some()
            && self.issue_year.unwrap() >= 2010
            && self.issue_year.unwrap() <= 2020
            // Expiration year
            && self.expiration_year.is_some()
            && self.expiration_year.unwrap() >= 2020
            && self.expiration_year.unwrap() <= 2030
            // Height
            && self.height.is_some()
            && self.height.as_ref().unwrap().is_valid()
            // Hair color
            && self.hair_color.is_some()
            && self.hair_color.as_ref().unwrap().starts_with("#")
            && self.hair_color.as_ref().unwrap().len() == 7
            // Eye color
            && self.eye_color.is_some()
            && eye_colors.contains(&self.eye_color.as_ref().unwrap().as_str())
            // Passport id
            && self.passport_id.is_some()
            && self.passport_id.as_ref().unwrap().len() == 9
    }
}

fn problem(content: &Vec<String>, should_validate: bool) -> Result<usize> {
    Ok(content
        .iter()
        .map(|line| Passport::try_from(line.to_owned()).unwrap())
        .filter(|passport| passport.has_values())
        .filter(|passport| {
            if should_validate {
                passport.is_valid()
            } else {
                true
            }
        })
        .count())
}

fn read_data() -> Vec<String> {
    include_str!("data.txt")
        .split("\n\n")
        .map(|s| s.replace("\n", " "))
        .collect()
}

fn main() -> Result<()> {
    let data = read_data();
    println!("Part 1: valid passports: {}", problem(&data, false)?);
    println!("Part 2: valid passports: {}", problem(&data, true)?);
    Ok(())
}

#[test]
fn test_part1() {
    let data = read_data();
    assert_eq!(problem(&data, false).unwrap(), 216);
}

#[test]
fn test_part2() {
    let data = read_data();
    assert_eq!(problem(&data, true).unwrap(), 150);
}
