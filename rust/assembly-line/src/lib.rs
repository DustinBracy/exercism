// This stub file contains items that aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

pub fn production_rate_per_hour(speed: u8) -> f64 {
    const MIN_RATE: f64 = 221.0;
    let production: f64 = speed as f64 * MIN_RATE;

    if speed > 8 {
        return production * 0.77
    }
    if speed > 4 {
        return production * 0.9
    }
    return production
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    return production_rate_per_hour(speed) as u32 / 60 
}
