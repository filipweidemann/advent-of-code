mod part1;

fn main() {
    const INPUT: &str = include_str!("input-1.txt");
    part1::get_total_score(INPUT);
    part1::get_total_score_part_two(INPUT);
}
