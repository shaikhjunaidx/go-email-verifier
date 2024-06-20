
# Email Domain Verifier

This Go program verifies email addresses by checking the associated domain's DNS records, including MX, SPF, and DMARC records. It can be useful for ensuring email deliverability and security by validating that the domain has the proper configurations for handling email.

## How It Works

1. **Input Reading:**
   - The program reads domain names from the standard input using `bufio.NewScanner(os.Stdin)`.

2. **DNS Record Checks:**
   - **MX Records:** Checks if the domain has Mail Exchange (MX) records, which are necessary for receiving emails.
   - **SPF Records:** Checks for the Sender Policy Framework (SPF) records to verify if the domain is authorized to send emails.
   - **DMARC Records:** Checks for Domain-based Message Authentication, Reporting, and Conformance (DMARC) records to ensure the domain has policies for email validation, reporting, and conformance.

3. **Output:**
   - The program prints the domain name along with boolean flags indicating the presence of MX, SPF, and DMARC records, and the actual SPF and DMARC records if present.

## Terms Used

- **MX (Mail Exchange) Records:**
  - These records specify the mail servers responsible for receiving email on behalf of a domain. They are essential for directing email traffic to the correct servers.

- **SPF (Sender Policy Framework) Records:**
  - SPF records are a type of TXT record that lists the IP addresses authorized to send email on behalf of a domain. They help prevent email spoofing by allowing the receiving mail server to verify the sender's IP address.

- **DMARC (Domain-based Message Authentication, Reporting, and Conformance) Records:**
  - DMARC records are TXT records that work with SPF and DKIM (DomainKeys Identified Mail) to provide a way for email receivers to authenticate emails from a domain. They help in reducing email fraud and phishing.

## Example Usage

To use this program, simply run it and provide domain names as input. For example:

```sh
go run main.go
example.com
anotherdomain.com
```

The program will output:

```
domain | hasMX | hasSPF | spfRecord | hasDMARC | dmarcRecord
example.com | true | true | v=spf1 include:_spf.example.com ~all | true | v=DMARC1; p=none; rua=mailto:dmarc-reports@example.com
anotherdomain.com | false | false |  | false |
```

## Potential Improvements

1. **Email Format Validation:**
   - Enhance the program to accept email addresses, validate their format, and extract domains automatically.
   
2. **Retry Logic:**
   - Implement retry logic for DNS lookups to handle transient network issues.

3. **SMTP Verification:**
   - Add functionality to perform SMTP verification by attempting to establish a connection to the mail server to further validate email addresses.

4. **Error Handling:**
   - Improve error handling to distinguish between different types of DNS lookup failures.

## Applications

- **Email Verification Services:**
  - Can be used in services that verify email addresses during user registration or form submissions to ensure valid and deliverable email addresses.

- **Email Security Solutions:**
  - Useful for security tools that monitor and validate email configurations to prevent email spoofing and phishing attacks.

- **Domain Configuration Audits:**
  - Can be part of auditing tools to check the email-related DNS configurations of a domain and ensure compliance with best practices.

**Next Steps:**
 - Add unit tests to validate the domain checking functionality.
 - Enhance the program to accept and validate email addresses directly.
