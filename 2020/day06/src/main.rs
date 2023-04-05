use std::collections::HashSet;

use itertools::Itertools;

fn part1() -> usize {
    include_str!("data.txt")
        .split("\n\n")
        .map(|group| group.replace("\n", ""))
        .map(|line| line.chars().unique().count())
        .sum()
}

fn part2() -> usize {
    include_str!("data.txt")
        .split("\n\n")
        .map(|group| {
            group
                .lines()
                .map(|person| person.chars().collect::<HashSet<char>>())
                .reduce(|a, b| &a & &b)
                .unwrap_or_default()
                .len()
        })
        .sum()
}

fn main() {
    println!("Part 1: {}", part1());
    println!("Part 2: {}", part2());
}

#[test]
fn test_part1() {
    assert_eq!(part1(), 6630);
}

#[test]
fn test_part2() {
    assert_eq!(part2(), 3437);
}
