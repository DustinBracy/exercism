codon_dict = {
    "AUG": "Methionine",
    "UUU": "Phenylalanine",
    "UUC": "Phenylalanine",
    "UUA": "Leucine",
    "UUG": "Leucine",
    "UCU": "Serine",
    "UCC": "Serine",
    "UCA": "Serine",
    "UCG": "Serine",
    "UAU": "Tyrosine",
    "UAC": "Tyrosine",
    "UGU": "Cysteine",
    "UGC": "Cysteine",
    "UGG": "Tryptophan",
}
stop_list = ["UAA", "UAG", "UGA"]


def proteins(strand: str) -> list:
    """Returns a list of protein names from a given RNA sequence."""
    protein_list = []
    while len(strand) >= 3:
        codon = strand[:3]
        if codon in stop_list:
            break
        strand = strand[3:]
        protein_list.append(codon_dict.get(codon))
    return protein_list
