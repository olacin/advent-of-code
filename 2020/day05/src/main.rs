use itertools::Itertools;

fn read_data() -> Vec<String> {
    include_str!("data.txt")
        .lines()
        .map(|line| format!("{}", line))
        .collect()
}

fn get_bounds(c: char, min: f32, high: f32) -> (f32, f32) {
    let delta = high - min;
    match c {
        'F' | 'L' => (min, (delta / 2.0).floor() + min),
        'B' | 'R' => ((delta / 2.0).ceil() + min, high),
        _ => (0.0, 0.0),
    }
}

fn process(s: &str) -> i32 {
    let mut row_bounds = (0.0, 127.0);

    // Find row
    for chr in s[..s.len() - 3].chars() {
        row_bounds = get_bounds(chr, row_bounds.0, row_bounds.1);
    }
    let row = row_bounds.0 as i32;

    let mut col_bounds = (0.0, 7.0);
    // Find column
    for chr in s[s.len() - 3..].chars() {
        col_bounds = get_bounds(chr, col_bounds.0, col_bounds.1);
    }
    let col = col_bounds.0 as i32;

    // Seat ID formula
    return row * 8 + col;
}

fn part1(data: &Vec<String>) -> i32 {
    data.iter().map(|line| process(line)).max().unwrap()
}

fn part2(data: &Vec<String>) -> i32 {
    // Sort seats
    let mut seats: Vec<i32> = data.iter().map(|line| process(line)).collect();
    seats.sort();
    // Find gap
    for seat in seats.iter().tuple_windows::<(_, _)>() {
        if (seat.1 - seat.0) > 1 {
            return seat.0 + 1;
        }
    }
    return -1;
}

fn main() {
    let data = read_data();
    // Part 1
    let max = part1(&data);
    println!("Highest seat ID: {}", max);
    // Part 2
    let seat = part2(&data);
    println!("My seat ID is {}", seat);
}

#[test]
fn test_part1() {
    let data = read_data();
    assert_eq!(part1(&data), 874);
}

#[test]
fn test_part2() {
    let data = read_data();
    assert_eq!(part2(&data), 594);
}
