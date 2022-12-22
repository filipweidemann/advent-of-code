use std::str::SplitWhitespace;

#[derive(Debug)]
struct Instruction {
    count: usize,
    source: usize,
    target: usize
}

pub enum CrateMover {
    Model9000,
    Model9001
}

fn parse_crates(input: &str) -> Vec<Vec<char>> {
    let mut crates = Vec::<Vec<char>>::new();

    input
        .lines()
        .rev()
        .skip(1)
        .for_each(|line| {
            line.chars()
                .skip(1)
                .step_by(4)
                .enumerate()
                .filter(|(_, item)| !item.is_ascii_whitespace())
                .for_each(|(i, item)| {
                    if i >= crates.len() {
                        let mut stack_chars = Vec::new();
                        stack_chars.push(item);
                        crates.push(stack_chars);
                    } else {
                        crates[i].push(item);
                    }
                })
    });
    
    println!("{:?}", crates);
    return crates
}

fn parse_instructions(instructions: &str) -> Vec<Instruction> {
    let mut parsed_instructions: Vec<Instruction> = Vec::new();

    for line in instructions.lines() {
        let instruction_numbers = line
            .split(" ")
            .skip(1)
            .step_by(2)
            .map(|i| i.parse::<usize>().unwrap())
            .collect::<Vec<usize>>();
        
        match instruction_numbers[..] {
            [count, source, target] => parsed_instructions.push(Instruction { count, source, target }),
            _ => ()
        }
    }
    
    println!("{:?}", parsed_instructions);
    return parsed_instructions
}

fn perform_instruction(instruction: &Instruction, crates: &mut Vec<Vec<char>>, cratemover_model: &CrateMover) {
    let source = crates.get_mut(instruction.source - 1).unwrap();

    // Instead of copying everything one-by-one, reverse the whole selection. Same outcome.
    let moved_crates = source.split_off(source.len() - instruction.count);
    
    let target = crates.get_mut(instruction.target - 1).unwrap();
    
    match cratemover_model {
        CrateMover::Model9000 => target.extend(moved_crates.iter().rev()),
        CrateMover::Model9001 => target.extend(moved_crates.iter())
    }
}

pub fn perform(input: &str, cratemover_model: CrateMover) -> (Vec<Vec<char>>, String) {
    let splitted: Vec<&str> = input.split("\n\n").collect();
    
    println!("SPLITTED: {:?}", splitted);
    // Maybe set max_stack_capacity longest stack * 2
    let mut crates = parse_crates(splitted[0]);
    let instructions = parse_instructions(splitted[1]);

    println!("{:?}", crates);
    println!("{:?}", instructions);

    for instruction in instructions {
        perform_instruction(&instruction, &mut crates, &cratemover_model);
    }
    
    println!("Moved Crates: {:?}", crates);
    
    // Get last items in order
    let mut order = String::from("");
    crates.iter().for_each(|crate_stack| {
        if let Some(c) = crate_stack.iter().last() {
            order.push(*c);
        }
    });

    println!("First crates: {:?}", order);
    return (crates, order)
}

#[cfg(test)]
mod tests {
    use crate::lib::{parse_crates, parse_instructions, perform};

    #[test]
    fn test_parse_crates() {
        const INPUT: &str = include_str!("input-test.txt");
        let crate_input = INPUT.split("\n\n").collect::<Vec<&str>>()[0];
        let crates = parse_crates(crate_input);
        assert_eq!(crates, vec![vec!['Z', 'N'], vec!['M', 'C', 'D'], vec!['P']]);
        // man. can  this be done in a sane way? shit macro...
    }

    #[test]
    fn test_parse_intruction() {
        const INPUT: &str = include_str!("input-test.txt");
        let instruction_input = INPUT.split("\n\n").collect::<Vec<&str>>()[1];
        let instructions = parse_instructions(instruction_input);
    }

    #[test]
    fn test_cratemover9000() {
        const INPUT: &str = include_str!("input-test.txt");
        let (_, order) = perform(INPUT, crate::lib::CrateMover::Model9000);
        assert_eq!(order, "CMZ");
    }

    #[test]
    fn test_cratemover9001() {
        const INPUT: &str = include_str!("input-test.txt");
        let (_, order) = perform(INPUT, crate::lib::CrateMover::Model9001);
        assert_eq!(order, "MCD");
    }
}
