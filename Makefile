# Coincheck MCP - Makefile
# Simple build and deployment automation

# Configuration
BINARY_NAME = coincheck-mcp
BUILD_DIR = ./build
INSTALL_DIR = ~/.local/bin
CONFIG_DIR = ~/.config/Claude
SOURCE_DIR = $(PWD)

# Go build settings
GOCMD = go
GOBUILD = $(GOCMD) build
GOMOD = $(GOCMD) mod
GOCLEAN = $(GOCMD) clean

# Default target
.PHONY: all
all: build install

# Show help
.PHONY: help
help:
	@echo "🔧 Coincheck MCP - Available commands:"
	@echo ""
	@echo "  make build     - Build the binary"
	@echo "  make install   - Install to ~/.local/bin"
	@echo "  make setup     - Setup configuration files"
	@echo "  make test      - Test the binary"
	@echo "  make clean     - Clean build artifacts"
	@echo "  make all       - Build and install (default)"
	@echo "  make deps      - Download dependencies"
	@echo "  make run       - Run with example config"
	@echo "  make help      - Show this help"
	@echo ""

# Download dependencies
.PHONY: deps
deps:
	@echo "📦 Downloading dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

# Build the binary
.PHONY: build
build: deps
	@echo "🔨 Building $(BINARY_NAME)..."
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) main.go
	@echo "✅ Build completed: $(BINARY_NAME)"

# Install to ~/.local/bin
.PHONY: install
install: build
	@echo "📁 Installing to $(INSTALL_DIR)..."
	@mkdir -p $(INSTALL_DIR)
	@cp $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_DIR)/
	@chmod +x $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "✅ Installed: $(INSTALL_DIR)/$(BINARY_NAME)"
	@echo "💡 You can now use: $(BINARY_NAME)"

# Setup configuration files
.PHONY: setup
setup:
	@echo "⚙️  Setting up configuration files..."
	@if [ ! -f config/coincheck_config.json ]; then \
		echo "📝 Creating coincheck_config.json from template..."; \
		cp config/coincheck_config.json.example config/coincheck_config.json; \
		echo "⚠️  Please edit config/coincheck_config.json with your API credentials"; \
	else \
		echo "✅ config/coincheck_config.json already exists"; \
	fi
	@echo "📝 Example config: config/claude_desktop_config.json.example"
	@echo "🎯 Next steps:"
	@echo "   1. Edit config/coincheck_config.json with your API credentials"
	@echo "   2. Use 'make deploy' to install MCP server to Claude Desktop"
	@echo "   3. Or configure Claude Desktop directly"

# Test the binary
.PHONY: test
test: install
	@echo "🧪 Testing $(BINARY_NAME)..."
	@if command -v $(BINARY_NAME) >/dev/null 2>&1; then \
		echo "✅ Binary is accessible in PATH"; \
		$(BINARY_NAME) --help | head -3; \
	else \
		echo "❌ Binary not found in PATH"; \
		echo "💡 Try: export PATH=$$PATH:$(INSTALL_DIR)"; \
		exit 1; \
	fi

# Run with example config (for testing)
.PHONY: run
run: build setup
	@echo "🚀 Running $(BINARY_NAME) with example config..."
	@if [ -f config/coincheck_config.json ]; then \
		$(BUILD_DIR)/$(BINARY_NAME) --config ./config/coincheck_config.json; \
	else \
		echo "❌ config/coincheck_config.json not found. Run 'make setup' first."; \
		exit 1; \
	fi

# Clean build artifacts
.PHONY: clean
clean:
	@echo "🧹 Cleaning build artifacts..."
	$(GOCLEAN)
	@rm -rf $(BUILD_DIR)
	@echo "✅ Clean completed"

# Force rebuild
.PHONY: rebuild
rebuild: clean all

# Development helpers
.PHONY: dev-setup
dev-setup: setup
	@echo "🛠️  Development setup..."
	@echo "📋 Configuration files created"
	@echo "💡 For development, you can use: make run"

# Show current configuration
.PHONY: status
status:
	@echo "📊 Project Status:"
	@echo "  Binary name: $(BINARY_NAME)"
	@echo "  Build dir:   $(BUILD_DIR)"
	@echo "  Install dir: $(INSTALL_DIR)"
	@echo "  Source dir:  $(SOURCE_DIR)"
	@echo ""
	@echo "📁 Files:"
	@ls -la $(BUILD_DIR)/$(BINARY_NAME) 2>/dev/null && echo "  ✅ Binary built" || echo "  ❌ Binary not built"
	@ls -la config/coincheck_config.json 2>/dev/null && echo "  ✅ Config exists" || echo "  ❌ Config not setup"
	@ls -la $(INSTALL_DIR)/$(BINARY_NAME) 2>/dev/null && echo "  ✅ Installed" || echo "  ❌ Not installed"

# Deploy to Claude Desktop (safe - no overwrite)
.PHONY: deploy
deploy: install
	@echo "🚀 Deploying to Claude Desktop..."
	@mkdir -p $(CONFIG_DIR)
	@if [ -f config/claude_desktop_config.json.example ]; then \
		if [ -f $(CONFIG_DIR)/claude_desktop_config.json ]; then \
			echo "⚠️  Existing Claude Desktop config found - skipping to prevent overwrite"; \
			echo "📝 To add coincheck MCP manually:"; \
			echo "   1. Open: $(CONFIG_DIR)/claude_desktop_config.json"; \
			echo "   2. Add coincheck entry from: config/claude_desktop_config.json.example"; \
			echo "   3. Restart Claude Desktop"; \
		else \
			cp config/claude_desktop_config.json.example $(CONFIG_DIR)/claude_desktop_config.json; \
			echo "✅ Config deployed to $(CONFIG_DIR)/"; \
			echo "🎯 Restart Claude Desktop to apply changes"; \
		fi \
	else \
		echo "❌ config/claude_desktop_config.json.example not found"; \
		exit 1; \
	fi
