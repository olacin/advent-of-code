fn problem(right: usize) -> usize {
    // TODO: make this work with down parameter
    let content: Vec<&[u8]> = include_str!("data.txt")
        .lines()
        .map(|line| line.as_bytes())
        .collect();

    let map_width = content[0].len();
    let trees = (0..content.len() - 1)
        .map(|idx| ((right * (idx + 1)) % map_width, idx + 1))
        .filter(|(x, y)| content[*y][*x] == b'#')
        .count();
    trees
}

fn part2() -> usize {
    let slopes = vec![1, 3, 5, 7, 1];
    slopes.iter().map(|slope| problem(*slope)).product()
}

fn main() {
    println!("{}", problem(3));
    println!("{}", part2());
}
