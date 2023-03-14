use itertools::Itertools;

fn problem(n: usize) -> i32 {
    include_str!("data.txt")
        .lines()
        .map(|s| s.parse::<i32>().unwrap())
        .combinations(n)
        .filter(|x| x.iter().sum::<i32>() == 2020)
        .map(|x| x.iter().product::<i32>())
        .next()
        .unwrap()
}

fn main() {
    let p1 = problem(2);
    println!("{p1}");
    let p2 = problem(3);
    println!("{p2}");
}

#[test]
fn test_part1() {
    assert_eq!(problem(2), 539851);
}

#[test]
fn test_part2() {
    assert_eq!(problem(3), 212481360);
}
