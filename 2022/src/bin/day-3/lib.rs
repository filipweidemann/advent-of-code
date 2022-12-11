use std::collections::HashSet;

struct Rucksack<'a> {
    contents: &'a str
}

#[derive(Debug)]
struct Compartment {
    contents: String
}

trait Storage {
    fn compartments(&self, number_of_compartments: &usize) -> Vec<Compartment>;
    fn duplicate(&self, number_of_compartments: &usize) -> Option<char>;
}

impl Storage for Rucksack<'_> {
    fn compartments(&self, number_of_compartments: &usize) -> Vec<Compartment> {
        return self.contents
            .chars()
            .collect::<Vec<char>>()
            .chunks((self.contents.chars().count()) / number_of_compartments)
            .map(|c| c.iter().collect::<String>())
            .map(|chunk| Compartment { contents: chunk })
            .collect::<Vec<Compartment>>();
    }

    fn duplicate(&self, number_of_compartments: &usize) -> Option<char> {
        let compartments: Vec<Compartment> = self.compartments(number_of_compartments);
        for character in compartments[0].contents.chars() {
            if compartments[1].contents.contains(character) {
                return Some(character)
            }
        }

        return None
    }
}

fn calculate_character_points(character: &char) -> usize {
    let uppercase: bool = character.is_uppercase();

    let alphabet: String = "abcdefghijklmnopqrstuvwxyz".to_string();
    let index = alphabet.find(character.to_lowercase().nth(0).unwrap());

    if let Some(index) = index {
        if !uppercase {
            return index + 1
        } else {
            return index + 27
        }
    }

    return 0;
}

fn find_common_in_groups<'a>(strings: &[&str]) -> usize {
    let s1 = strings[0];
    let s2 = strings[1];
    let s3 = strings[2];

    let charsets: Vec<HashSet<char>> = vec![
        HashSet::from_iter(s2.chars()),
        HashSet::from_iter(s3.chars())
    ];

    let common_char = s1
        .chars()
        .filter(|c| charsets[0].contains(&c) && charsets[1].contains(&c))
        .collect::<Vec<char>>();
    
    if let Some(common) = common_char.first() {
        calculate_character_points(common)
    } else {
        return 0
    }
}

pub fn calculate_item_priorities(input: &str) -> usize {
    return input.
        lines()
        .map(|line| Rucksack { contents: line })
        .map(|rucksack| rucksack.duplicate(&2))
        .map(|character| {
            if let Some(character) = character {
                return calculate_character_points(&character);
            } 

            return 0
        }).sum();
}

pub fn calculate_elve_group_priorities(input: &str) -> usize {
    input
        .lines()
        .collect::<Vec<_>>()
        .chunks_exact(3)
        .map(find_common_in_groups)
        .sum()
}


#[cfg(test)]
mod tests {
    use crate::lib::{calculate_item_priorities, calculate_elve_group_priorities};

    use super::{Rucksack, Storage, calculate_character_points};

    #[test]
    fn test_compartment_splitting() {
        let rucksack = Rucksack { contents: "ABFDEF" };
        let compartments = rucksack.compartments(&2);
        assert_eq!(compartments[0].contents, "ABF");
        assert_eq!(compartments[1].contents, "DEF");
    }
    
    #[test]
    fn test_duplication_detection() {
        let rucksack = Rucksack { contents: "ABFDEF" };
        assert_eq!(rucksack.duplicate(&2), Some('F'));
    }

    #[test]
    fn test_character_points_calculation() {
        let rucksack = Rucksack { contents: "abAa" };
        let rucksack_2 = Rucksack { contents: "AbaA" };
        let duplicate = rucksack.duplicate(&2);
        let duplicate_2 = rucksack_2.duplicate(&2);

        if let Some(duplicate_char) = duplicate {
            let score = calculate_character_points(&duplicate_char);
            assert_eq!(score, 1);
        }

        if let Some(duplicate_char) = duplicate_2 {
            let score = calculate_character_points(&duplicate_char);
            assert_eq!(score, 27)
        }
    }

    #[test]
    fn test_input_transformation() {
        const input: &str = include_str!("input-test.txt");
        let sum = calculate_item_priorities(input);
        assert_eq!(sum, 157)
    }

    #[test]
    fn test_elve_groups() {
        const input: &str = include_str!("input-test.txt");
        let sum = calculate_elve_group_priorities(input);
        assert_eq!(sum, 70);
    }
}
