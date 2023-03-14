use anyhow::Result;
use regex::Regex;

fn problem(validation_fn: &dyn Fn(usize, usize, &char, &str) -> bool) -> Result<usize> {
    let pattern = Regex::new(r"(\d+)-(\d+) ([a-z]): (.*)").unwrap();
    let count = include_str!("data.txt")
        .lines()
        .map(|line| {
            let captures = pattern.captures(line).unwrap();

            let start: usize = captures.get(1).unwrap().as_str().parse().unwrap();
            let end: usize = captures.get(2).unwrap().as_str().parse().unwrap();
            let letter = captures.get(3).unwrap().as_str().chars().next().unwrap();
            let pwd = captures.get(4).unwrap().as_str();

            return (start, end, letter, pwd);
        })
        .filter(|(start, end, letter, pwd)| validation_fn(*start, *end, letter, pwd))
        .count();
    Ok(count)
}

fn policy_part1(start: usize, end: usize, letter: &char, pwd: &str) -> bool {
    let count = pwd.chars().filter(|c| c == letter).count();
    count >= start && count <= end
}

fn policy_part2(start: usize, end: usize, letter: &char, pwd: &str) -> bool {
    let p1 = pwd.chars().nth(start - 1).unwrap();
    let p2 = pwd.chars().nth(end - 1).unwrap();
    (p1 == *letter) ^ (p2 == *letter)
}

fn main() -> Result<()> {
    let count = problem(&policy_part1)?;
    println!("{count}");
    let count2 = problem(&policy_part2)?;
    println!("{count2}");
    Ok(())
}

#[test]
fn test_part1() {
    assert_eq!(problem(&policy_part1).unwrap(), 416);
}

#[test]
fn test_part2() {
    assert_eq!(problem(&policy_part2).unwrap(), 688);
}
