document.addEventListener("DOMContentLoaded", function () {
    const questionText = document.getElementById("question-text");
    const option1Label = document.getElementById("label-option1");
    const option2Label = document.getElementById("label-option2");
    const option3Label = document.getElementById("label-option3");
    const option4Label = document.getElementById("label-option4");
    const nextButton = document.getElementById("next-button");
    const mcqForm = document.getElementById("mcq-form");
    let currentQuestionIndex = 0;

    // Function to fetch questions from the server
    async function fetchQuestions() {
        try {
            const response = await fetch("/home/java/mcq"); // Adjust the URL accordingly
            if (!response.ok) {
                throw new Error("Failed to fetch questions.");
            }
            const questions = await response.json();
            return questions;
        } catch (error) {
            console.error(error);
            alert("Failed to fetch questions.");
            return [];
        }
    }

    // Function to display the current question
    async function displayQuestion(index) {
        const questions = await fetchQuestions();
        if (index < questions.length) {
            const currentQuestion = questions[index];
            questionText.textContent = currentQuestion.MQ_question;
            option1Label.textContent = currentQuestion.Option1;
            option2Label.textContent = currentQuestion.Option2;
            option3Label.textContent = currentQuestion.Option3;
            option4Label.textContent = currentQuestion.Option4;
        }
    }

    // Initial display of the first question
    displayQuestion(currentQuestionIndex);

    // Function to check the answer and go to the next question
    function checkAnswerAndNext() {
        // Implement your answer checking logic here

        // Move to the next question
        currentQuestionIndex++;
        displayQuestion(currentQuestionIndex);

        // Clear the selected answer
        const selectedAnswer = document.querySelector('input[name="answer"]:checked');
        if (selectedAnswer) {
            selectedAnswer.checked = false;
        }

        if (currentQuestionIndex >= questions.length) {
            // All questions answered, you can add your logic here
            alert("All questions answered!");
        }
    }

    // Add a click event listener to the Next button
    nextButton.addEventListener("click", checkAnswerAndNext);
});
