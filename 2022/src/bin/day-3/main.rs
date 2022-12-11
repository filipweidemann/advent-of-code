mod lib;

fn main() {
    const INPUT_STR: &str = include_str!("input-1.txt");
    let sum = lib::calculate_item_priorities(INPUT_STR);
    println!("Rucksack Priority: {}", sum);

    let sum2 = lib::calculate_elve_group_priorities(INPUT_STR);
    println!("Elve Group Priority: {}", sum2)
}
