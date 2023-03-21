fn read_data() -> Vec<&'static [u8]> {
    include_str!("data.txt")
        .lines()
        .map(|line| line.as_bytes())
        .collect()
}

fn problem(content: &Vec<&[u8]>, right: usize, down: usize) -> usize {
    let map_width = content[0].len();
    content
        .iter()
        .step_by(down)
        .enumerate()
        .filter(|(idx, row)| row[idx * right % map_width] == b'#')
        .count()
}

fn part2(data: &Vec<&[u8]>) -> usize {
    vec![(1, 1), (3, 1), (5, 1), (7, 1), (1, 2)]
        .iter()
        .map(|(right, down)| problem(data, *right, *down))
        .product()
}

fn main() {
    let data = read_data();
    println!("Part 1: {}", problem(&data, 3, 1));
    println!("Part 2: {}", part2(&data));
}

#[test]
fn test_part1() {
    let data = read_data();
    assert_eq!(problem(&data, 3, 1), 284);
}

#[test]
fn test_part2() {
    let data = read_data();
    assert_eq!(part2(&data), 3510149120);
}
