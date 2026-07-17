object Bob {
    fun hey(input: String): String {
        var text = input.toString().trim()
        val hasChars = text.any { it.isLetter() }
        val isAllCaps = text == text.uppercase()

        val output = when {
            text == "" -> "Fine. Be that way!"
            isAllCaps && hasChars && text.endsWith("?") -> "Calm down, I know what I'm doing!"
            isAllCaps && hasChars -> "Whoa, chill out!"
            text.endsWith("?") -> "Sure."
            else -> "Whatever."
        }

        return output
    }
}
