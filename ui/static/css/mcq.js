document.addEventListener("DOMContentLoaded", function() {
    var options = document.querySelectorAll('.option');

    options.forEach(function(option, index) {
        option.addEventListener("click", function() {
            handleOptionClick(option, index);
        });
    });

    function handleOptionClick(option, index) {
        var radioInput = option.querySelector('input[type="radio"]');

        if (radioInput.checked) {
            return;
        }

        if (parseInt(index) === parseInt(document.getElementById("correct-answer").value)) {
            option.classList.remove('wrong');
            option.classList.add('correct');
        } else {
            option.classList.remove('correct');
            option.classList.add('wrong');

            options.forEach(function(opt) {
                if (parseInt(opt.querySelector('input[type="radio"]').value) === parseInt(document.getElementById("correct-answer").value)) {
                    opt.classList.remove('wrong');
                    opt.classList.add('correct');
                }
            });
        }

        var radioInputs = document.querySelectorAll('input[type="radio"]');
        radioInputs.forEach(function(input) {
            if (input !== radioInput) {
                input.checked = false;
            }
        });

        radioInput.checked = true;

        fetch("/check-answer", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                isCorrect: parseInt(index) === parseInt(document.getElementById("correct-answer").value),
            }),
        })
        .then((response) => response.json())
        .then((data) => {
            console.log("Server response:", data);
        })
        .catch((error) => {
            console.error("Error sending request to the server:", error);
        });
    }
});
