mod lib;

fn main() {
    const INPUT_STR: &str = include_str!("input.txt");
    
    // Part 1
    println!("Index of last match sequence (Signal): {:?}", lib::find_signal_marker(INPUT_STR, &lib::Marker::Signal));

    // Part 2
    println!("Index of last match sequence (Message): {:?}", lib::find_signal_marker(INPUT_STR, &lib::Marker::Message));
}
