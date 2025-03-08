# Calculator

Calculator is a Command Line Interface (CLI) calculator written in Go.

## Installation/Usage

To run the calculator, you can use GitHub Actions to build and test the project without needing to install Go on your local machine.

### Using GitHub Actions for CI/CD

This project uses GitHub Actions for Continuous Integration and Continuous Deployment (CI/CD).

### Workflow Details

The workflow is defined in the `.github/workflows/go.yml` file and includes the following steps:
- Checkout code
- Set up Go environment
- Install dependencies
- Build the project
- Run tests
- Upload build artifacts

### Running the CI/CD Workflow

The CI/CD workflow runs automatically on `push` and `pull_request` events to the `main` branch. You can view the workflow status and logs in the "Actions" tab of the repository.

### Downloading Built Artifacts

1. **Navigate to the Actions Tab:**
   Go to the "Actions" tab in the GitHub repository.

2. **Select the Workflow Run:**
   Click on the workflow run you are interested in.

3. **Download the Artifacts:**
   Locate the "Upload build artifacts" step and click on the link to download the compiled binary.

## Contribution

If you would like to contribute to this project, please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License.
