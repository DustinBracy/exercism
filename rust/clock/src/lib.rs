use std::fmt;

#[derive(Debug)]
pub struct Clock {
    hours: i32,
    minutes: i32,
}

impl Clock {
    pub fn add_minutes(&mut self, minutes: i32) -> Self {
        self.minutes += minutes;
        self.normalize()
    }

    pub fn normalize(&mut self) -> Self {
        let add_hours = (self.minutes) / 60;
        let total_hours = add_hours + self.hours;
        let mut new_hours;
        let new_minutes;

        match self.minutes % 60 {
            m if m < 0 => {
                new_hours = total_hours - 1;
                new_minutes = m + 60;
            }
            m if m > 59 => {
                new_hours = total_hours + 1;
                new_minutes = m - 60
            }
            m => {
                new_hours = total_hours;
                new_minutes = m % 60;
            }
        }

        if new_hours < 0 {
            new_hours = new_hours % 24 + 24
        }

        Self {
            hours: new_hours % 24,
            minutes: new_minutes,
        }
    }

    pub fn new(hours: i32, minutes: i32) -> Self {
        let mut new_clock = Self { hours, minutes };
        new_clock.normalize()
    }
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{:02}:{:02}", self.hours, self.minutes)
    }
}

impl PartialEq for Clock {
    fn eq(&self, other: &Clock) -> bool {
        format!("{}", self) == format!("{}", other)
    }
}
