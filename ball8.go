package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    answers := []string{
        "Definitely.",
        "Most likely yes.",
        "Yes, but you have to be careful.",
        "I'm not sure, please try again.",
        "Don't count on it.",
        "Definitely not.",
        "The answer is unclear; please try again.",
        "I would like to, but I can't.",
        "Definitely YES.",
        "I'm for it, and even against it.",
        "Keep an open mind.",
        "Follow the right path.",
        "Unlock your potential.",
        "If you're ready for it.",
        "It depends on your inner experience.",
        "Keep seeking further answers.",
        "Identify and solve your internal blocks.",
        "Follow your intuitions.",
        "Trust yourself.",
        "Be open.",
        "It's possible, but I'm not sure.",
        "Let intuition guide you to the answer.",
        "Maybe it's worth discussing with loved ones.",
        "The answer lies within you, you just have to search for it.",
        "There is no definitive answer to this question.",
        "Maybe it's worth asking yourself this question again in the future.",
        "Don't close yourself off to different perspectives, the answer may be more complex than it seems.",
        "This question requires deeper considerations.",
        "The answer depends on the circumstances and choices you make.",
        "Don't rush with the answer, time will show what is right for you.",
        "Definitely yes!",
        "Without a doubt.",
        "Of course!",
        "Absolutely.",
        "No way.",
        "Probably not.",
        "Focus and try again.",
        "Don't count on it, but who knows...",
        "I doubt it.",
        "We'll see.",
        "I don't know the answer to this question.",
        "Maybe.",
        "I don't say this often, but yes.",
        "The answer is in your hands.",
        "I can't predict right now.",
        "Probably not.",
        "Don't count on it anytime soon.",
        "Signs point to yes.",
        "Maybe it's better not to ask.",
        "It answer definitively.",
        "I can't answer definitively.",
        "Perhaps.",
        "I don't think so.",
        "Probably yes, but I'm not sure.",
        "Focus on your goals, and the answer will become clear.",
        "It's hard to say.",
        "Don't count on it anytime soon.",
        "I can't answer that question.",
        "The answer lies somewhere in the middle.",
        "Forget about this question.",
        "I wouldn't be certain.",
        "You can count on it.",
        "Don't do it.",
        "The answer is so obvious that there shouldn't be a question.",
        "Try your luck another time.",
        "Don't expect it.",
        "Focus on more important matters.",
        "You shouldn't count on it.",
        "The answer is contrary to your expectations.",
        "This is a question only you know the answer to.",
        "Don't lose hope.",
        "Probably yes, but I'm not sure.",
        "There's no answer to this question.",
        "Focus on something else.",
        "I wouldn't have a clue.",
        "Try asking someone else.",
        "The answer is on the horizon.",
    }

    rand.Seed(time.Now().UnixNano())

    var question string
    fmt.Println("Enter a question:")
    fmt.Scanln(&question)

    answer := answers[rand.Intn(len(answers))]

    fmt.Println("Ball8 speak:", answer)
}
