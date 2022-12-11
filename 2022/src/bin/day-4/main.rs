
mod lib;

fn main() {
    const INPUT_STR: &str = include_str!("input.txt");
    println!("Full overlaps: {:?}", lib::count_full_overlaps(INPUT_STR));
    println!("Partial overlaps: {:?}", lib::count_partial_overlaps(INPUT_STR));
}
