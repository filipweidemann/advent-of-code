mod lib;

fn main() {
    const INPUT_STR: &str = include_str!("input.txt");
    
    // Part 1
    let (_, order) = lib::perform(INPUT_STR, lib::CrateMover::Model9000);
    println!("{:?}", order);

    // Part 2
    let (_, order) = lib::perform(INPUT_STR, lib::CrateMover::Model9001);
    println!("{:?}", order);
}
