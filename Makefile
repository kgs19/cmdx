# 	Add go fmt to format the code
fmt:
	@echo "Formatting code..."
	go fmt
	@echo "Done!"

run-test:
	@echo "Running tests..."
	go test -v
	@echo "Done!"

run-examples:
	@echo "Running examples..."
	cd examples && \
	go run run_ps_print_output.go && \
	go run run_date_return_output.go && \
	go run run_ls_print_output_with_dir.go && \
	go run run_ls_custom_config.go
	@echo "Done!"


