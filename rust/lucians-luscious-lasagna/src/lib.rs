// This stub file contains items that aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

use core::num;

const EXPECTED_MINUTES_IN_OVEN: i32 = 40;
const LAYER_PREPERATION_TIME_MINUTES: i32 = 2;

pub fn expected_minutes_in_oven() -> i32 {
    return EXPECTED_MINUTES_IN_OVEN;
}

pub fn remaining_minutes_in_oven(actual_minutes_in_oven: i32) -> i32 {
    return EXPECTED_MINUTES_IN_OVEN - actual_minutes_in_oven;
}

pub fn preparation_time_in_minutes(number_of_layers: i32) -> i32 {
    return LAYER_PREPERATION_TIME_MINUTES * number_of_layers;
}

pub fn elapsed_time_in_minutes(number_of_layers: i32, actual_minutes_in_oven: i32) -> i32 {
    return preparation_time_in_minutes(number_of_layers) + actual_minutes_in_oven;
}
