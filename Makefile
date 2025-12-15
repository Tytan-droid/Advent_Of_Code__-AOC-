runall:
	@if [ -z "$(year)" ]; then \
		echo "Usage: make runall year=<year>"; exit 1; \
	fi
	@for d in $$(ls -d $(year)/D* 2>/dev/null | sort -V); do \
		echo "▶ Exécution $$d"; \
		go run $$d/solution.go; \
		echo "\n\n\n"; \
	done

run:
	@if [ -z "$(year)" ] || [ -z "$(day)" ]; then \
		echo "Usage: make run year=<year> day=<numero>"; exit 1; \
	fi
	go run $(year)/D$(day)/solution.go