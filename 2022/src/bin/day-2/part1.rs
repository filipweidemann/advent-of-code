use std::collections::HashMap;

fn points_mapping() -> HashMap<&'static str, u32> {
    let mut strategy_points: HashMap<&str, u32> = HashMap::new();
    // "Legacy" Scoring
    strategy_points.insert("A X", 4);
    strategy_points.insert("A Y", 8);
    strategy_points.insert("A Z", 3);
    strategy_points.insert("B Y", 5);
    strategy_points.insert("B X", 1);
    strategy_points.insert("B Z", 9);
    strategy_points.insert("C Z", 6);
    strategy_points.insert("C Y", 2);
    strategy_points.insert("C X", 7);
    
    // Clarified Scoring
    strategy_points.insert("A A", 4);
    strategy_points.insert("A B", 8);
    strategy_points.insert("A C", 3);
    strategy_points.insert("B B", 5);
    strategy_points.insert("B A", 1);
    strategy_points.insert("B C", 9);
    strategy_points.insert("C C", 6);
    strategy_points.insert("C B", 2);
    strategy_points.insert("C A", 7);
    return strategy_points
}

fn get_points(play_move: &str) -> u32 {
    let points_mapping = points_mapping();
    let result = points_mapping.get(play_move);

    if let Some(result) = result {
        return *result;
    }

    return 1
}

struct PlayOption<'a> {
    name: &'a str,
    beats: &'a str,
    beaten_by: &'a str
}

fn possible_moves() -> Vec<PlayOption<'static>> {
    let mut options = Vec::new();
    options.push(PlayOption { name: "A", beats: "C", beaten_by: "B" });
    options.push(PlayOption { name: "A", beats: "C", beaten_by: "B" });
    options.push(PlayOption { name: "B", beats: "A", beaten_by: "C" });
    options.push(PlayOption { name: "C", beats: "B", beaten_by: "A" });
    return options
}

fn get_required_play(opponent_move: &str, desired_outcome: & str) -> String {
    let required_move: Vec<PlayOption> = possible_moves().into_iter().filter(|possible_move| possible_move.name == opponent_move).collect();
    let required_move_option = required_move.get(0);
    
    if let Some(required_move_option) = required_move_option {
        if desired_outcome == "X" {
            return format!("{} {}", opponent_move, required_move_option.beats)
        } else if desired_outcome == "Z" {
            return format!("{} {}", opponent_move, required_move_option.beaten_by)
        } else {
            // Mirror Option to draw the game
            return format!("{} {}", opponent_move, opponent_move)
        }
    } else {
        panic!()
    } 
}

pub fn get_total_score(strategy_input: &str) {
    let score: u32 = strategy_input.lines().map(|line| get_points(line)).sum();
    println!("{}", score)
}

pub fn get_total_score_part_two(strategy_input: &str) {
    let score: u32 = strategy_input.lines().map(|line| {
        let splitted_lines = line.split(" ").collect::<Vec<&str>>();
        let required_play = get_required_play(splitted_lines[0], splitted_lines[1]);
        return required_play
    }).map(|required_play| get_points(&required_play)).sum();
    println!("{}", score);
}
